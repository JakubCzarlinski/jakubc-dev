package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	omitempty "github.com/JakubCzarlinski/go-add-omit-empty"
	"github.com/JakubCzarlinski/go-logging"
	go2ts "github.com/JakubCzarlinski/go-struct-to-ts"
	"golang.org/x/sync/errgroup"
)

const svelteDir string = "."

const renderToTemplDir string = "./build/render_to_templ/"
const svelteCompileDir string = "./compile/"
const databaseConfigDir string = "./project/"
const golangRootDir string = "./project/"
const genDir string = "./project/gen/"
const srcDir string = "./project/src/"
const sqlcOutDir string = "./project/src/data/"
const libDir string = "./project/src/lib/"
const genSqlc bool = false

var permittedOmitJsonFilenames = [...]string{"models.go"}

func runProcess(dir string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return logging.Bubble(err, "Error running process")
	}
	return nil
}

func main() {
	args := os.Args[1:]
	startTime := time.Now()
	var err error

	// If no arguments are passed, build the project in full
	if len(args) == 0 {
		logging.Info("Building project in full...")

		frontendBuildStartTime := time.Now()
		var golangToTypescriptElapsedTime time.Duration
		var bundlingElapsedTime time.Duration
		var renderElapsedTime time.Duration
		var sqlcCompileElapsedTime time.Duration
		var formattingElapsedTime time.Duration
		var generateElapsedTime time.Duration
		var serverBuildElapsedTime time.Duration
		var builderElapsedTime time.Duration
		var frontendBuildElapsedTime time.Duration
		var bundlingWaitGroup sync.WaitGroup
		bundlingWaitGroup.Add(1)
		go func() {
			defer bundlingWaitGroup.Done()
			bundlingStartTime := time.Now()
			err = viteBundle()
			if err != nil {
				logging.FatalF(err.Error())
			}
			bundlingElapsedTime = time.Since(bundlingStartTime)
		}()

		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			renderStartTime := time.Now()
			err = serverSideRender()
			if err != nil {
				logging.FatalF(err.Error())
			}
			renderElapsedTime = time.Since(renderStartTime)
		}()

		go func() {
			defer wg.Done()
			sqlcCompileStartTime := time.Now()
			err = sqlcGenerate()
			if err != nil {
				logging.FatalF(err.Error())
			}
			sqlcCompileElapsedTime = time.Since(sqlcCompileStartTime)
		}()

		wg.Wait()

		frontendBuildElapsedTime = time.Since(frontendBuildStartTime)

		builderStartTime := time.Now()
		err = renderToTempl()
		if err != nil {
			logging.FatalF(err.Error())
		}
		builderElapsedTime = time.Since(builderStartTime)

		formattingStartTime := time.Now()
		err = templFormat()
		if err != nil {
			logging.FatalF(err.Error())
		}
		formattingElapsedTime = time.Since(formattingStartTime)

		generateStartTime := time.Now()
		err = templGenerate()
		if err != nil {
			logging.FatalF(err.Error())
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			golangToTypescriptStartTime := time.Now()

			err = convertStructsToTypescript()
			if err != nil {
				logging.FatalF(err.Error())
			}

			golangToTypescriptElapsedTime = time.Since(golangToTypescriptStartTime)
		}()

		err = goFormat()
		if err != nil {
			logging.FatalF(err.Error())
		}

		generateElapsedTime = time.Since(generateStartTime)

		serverBuildStartTime := time.Now()
		err = buildGoServer()
		if err != nil {
			logging.FatalF(err.Error())
		}
		serverBuildElapsedTime = time.Since(serverBuildStartTime)

		bundlingWaitGroup.Wait()
		wg.Wait()

		elapsedTime := time.Since(startTime)

		fmt.Println(logging.Cyan("Waterfall diagram:"))
		fmt.Println(logging.Cyan(fmt.Sprintf("1---> Frontend build: %d ms", frontendBuildElapsedTime.Milliseconds())))
		fmt.Println(logging.Cyan(fmt.Sprintf("|---\\---1---> Go2TS: %d ms", golangToTypescriptElapsedTime.Milliseconds())))
		fmt.Println(logging.Cyan(fmt.Sprintf("|---\\---2---> Bundling: %d ms", bundlingElapsedTime.Milliseconds())))
		fmt.Println(logging.Cyan(fmt.Sprintf("|---|---> Rendering: %d ms", renderElapsedTime.Milliseconds())))
		fmt.Println(logging.Cyan(fmt.Sprintf("|---> SQLC compile: %d ms", sqlcCompileElapsedTime.Milliseconds())))
		fmt.Println(logging.Cyan(fmt.Sprintf("2---> Builder: %d ms", builderElapsedTime.Milliseconds())))
		fmt.Println(logging.Cyan(fmt.Sprintf("3---> Formatting: %d ms", formattingElapsedTime.Milliseconds())))
		fmt.Println(logging.Cyan(fmt.Sprintf("4---> Generate: %d ms", generateElapsedTime.Milliseconds())))
		fmt.Println(logging.Cyan(fmt.Sprintf("5---> Server build: %d ms", serverBuildElapsedTime.Milliseconds())))
		fmt.Println(logging.Cyan(fmt.Sprintf("- Total: %d ms", elapsedTime.Milliseconds())))
	}

	// If first argument is "go", build the Go server
	if len(args) > 0 && args[0] == "go" {
		buildStep := [](func() error){
			templFormat,
			templGenerate,
			goFormat,
			buildGoServer,
		}
		runBuild(buildStep)

	} else if len(args) > 0 && args[0] == "sql" {
		buildStep := [](func() error){
			sqlcGenerate,
			goFormat,
			buildGoServer,
		}
		runBuild(buildStep)
	}
}

func runBuild(buildStep []func() error) {
	var err error
	for _, step := range buildStep {
		err = step()
		if err != nil {
			logging.Fatal(err.Error())
		}
	}
}

func renderToTempl() error {
	err := runProcess(".", renderToTemplDir+"main.exe", "-in", svelteCompileDir, "-out", genDir)
	if err != nil {
		return logging.Bubble(err, "Error running frontend builder")
	}
	return nil
}

func templFormat() error {
	os.Setenv("TEMPL_EXPERIMENT", "rawgo")
	err := runProcess(golangRootDir, "templ", "fmt")
	if err != nil {
		return logging.Bubble(err, "Error formatting code")
	}
	return nil
}

func templGenerate() error {
	os.Setenv("TEMPL_EXPERIMENT", "rawgo")
	err := runProcess(golangRootDir, "templ", "generate")
	if err != nil {
		return logging.Bubble(err, "Error generating templates")
	}
	return nil
}

func goFormat() error {
	err := runProcess(golangRootDir, "gofmt", "-w", "./src/")
	if err != nil {
		return logging.Bubble(err, "Error running gofmt")
	}
	return nil
}

func buildGoServer() error {
	err := runProcess(golangRootDir, "go", "build", "-tags", "sonic avx", "-ldflags=-s -w", "-o", "main.exe", "./src/main.go")
	if err != nil {
		return logging.Bubble(err, "Error building Go server")
	}
	return nil
}

func sqlcGenerate() error {
	if !genSqlc {
		return nil
	}
	err := runProcess(databaseConfigDir, "sqlc", "generate")
	if err != nil {
		return logging.Bubble(err, "Error compiling SQL")
	}

	files, err := os.ReadDir(sqlcOutDir)
	if err != nil {
		return logging.Bubble(err, "Error reading directory for sqlc output")
	}

	sqlcWaitGroup := errgroup.Group{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".sql.go") {
			found := false
			for _, permitted := range permittedOmitJsonFilenames {
				if file.Name() == permitted {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}
		sqlcWaitGroup.Go(func() error {
			return omitempty.AddOmitJson(sqlcOutDir + file.Name())
		})
	}
	err = sqlcWaitGroup.Wait()
	if err != nil {
		return logging.Bubble(err, "Error adding omitempty to JSON tags")
	}
	return nil
}

func convertStructsToTypescript() error {
	innerWaitGroup := errgroup.Group{}

	currentDir := srcDir
	go2ts.SetRootDirName("project")
	err := filepath.WalkDir(srcDir, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		fileNameLen := len(info.Name())
		dir := path[:len(path)-fileNameLen]
		if dir == currentDir {
			return nil
		}
		currentDir = dir

		innerWaitGroup.Go(func() error {
			foundTypes, err := go2ts.ReadTypes(dir)
			if err != nil {
				return logging.Bubble(err, "Error reading types")
			}
			if foundTypes == "" {
				return nil
			}

			convertedTypes, err := go2ts.Convert(foundTypes)
			if err != nil {
				return logging.Bubble(err, "Error converting types")
			}
			if convertedTypes == "" {
				return nil
			}
			return os.WriteFile(dir+"types.go.ts", []byte(convertedTypes), 0644)
		})
		return nil
	})

	if err != nil {
		return logging.Bubble(err, "Error walking directory")
	}

	err = innerWaitGroup.Wait()
	if err != nil {
		return logging.Bubble(err, "Error converting Go structs to TypeScript")
	}
	return nil
}

func viteBundle() error {
	err := runProcess(svelteDir, "bun", "run", "--silent", "vite", "build", "--logLevel", "error")
	if err != nil {
		return logging.Bubble(err, "Error building Svelte frontend")
	}
	return nil
}

func serverSideRender() error {
	err := runProcess(svelteDir, "rm", "-rf", svelteCompileDir+"*")
	if err != nil {
		return logging.Bubble(err, "Error removing queue directory")
	}

	err = runProcess(svelteDir, "bun", "run", "--logLevel=error", "./build/render/dist/main.js", "-i", libDir, "-o", svelteCompileDir)
	if err != nil {
		return logging.Bubble(err, "Error running render process")
	}

	err = runProcess(svelteDir, "prettier", "--log-level", "error", "--write", "--single-attribute-per-line=false", "--bracket-same-line=true", svelteCompileDir+"**/*.html")
	if err != nil {
		return logging.Bubble(err, "Error running prettier")
	}
	return nil
}
