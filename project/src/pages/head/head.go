package head

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"project/src/flags"
	"strings"

	"github.com/JakubCzarlinski/go-logging"
	"github.com/JakubCzarlinski/go-pooling"
	"github.com/a-h/templ"
)

var indexScripts string
var stylesheets map[string]string = make(map[string]string)

var bufferPool = pooling.CreateBytesBufferPool(1024 * 12)

func init() {
	var ex, _ = os.Executable()
	var exPath = filepath.Dir(ex)
	file, err := os.Open(path.Join(exPath, flags.DistDir+"index.html"))
	if err != nil {
		logging.FatalBubble(err, "Error reading file")
	}
	defer file.Close()
	// Read the file into indexScripts
	headTags := &strings.Builder{}
	_, err = io.Copy(headTags, file)
	if err != nil {
		logging.FatalBubble(err, "Error reading file")
	}
	indexScripts = headTags.String()
	indexScripts = strings.ReplaceAll(indexScripts, "\n", "")

	files, err := os.ReadDir(path.Join(exPath, flags.AssestsDir))
	if err != nil {
		logging.FatalBubble(err, "Error reading directory")
	}

	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".css") {
			stylesheets[filename[:len(filename)-4]] = fmt.Sprintf(`<link rel="stylesheet" href="/assets/%s"/>`, filename)
		}
	}
}

func DefaultPageRender(
	contents templ.Component,
	headContents map[string]struct{},
	res http.ResponseWriter,
	lightMode bool,
) {
	ctx := context.Background()
	bodyBuffer := bufferPool.Get()
	defer bufferPool.Reset(bodyBuffer, struct{}{})

	// Render the contents first.
	bodyBuffer.WriteString(`<body data-theme="custom" class="overflow-x-hidden overflow-y-auto">`)
	err := contents.Render(ctx, bodyBuffer)
	if err != nil {
		panic(err)
	}
	bodyBuffer.WriteString("</body>")

	// Render the head contents.
	if lightMode {
		res.Write([]byte(`<!doctype html><html lang="en">`))
	} else {
		res.Write([]byte(`<!doctype html><html lang="en" class="dark">`))
	}

	err = head(headContents).Render(ctx, res)
	if err != nil {
		panic(err)
	}

	// Render the default page.
	res.Write(bodyBuffer.Bytes())
	res.Write([]byte("</html>"))
}

func LinkImgToHead(headContents map[string]struct{}, src string) {
	headContents[`<link rel="preload" as="image" href="`+src+`"/>`] = struct{}{}
}

func createHeadContents(headContents map[string]struct{}) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, indexScripts)
		if err != nil {
			return err
		}

		for key := range headContents {
			_, err = io.WriteString(w, key)
			if err != nil {
				return err
			}
			indexHref := strings.Index(key, `href="/assets/`)
			if indexHref == -1 {
				continue
			}
			remaining := key[indexHref+14:] // 14 is the length of `href="/assets/`
			jsIndex := strings.Index(remaining, `.js"`)
			if jsIndex == -1 {
				continue
			}

			filename := remaining[:jsIndex]
			stylesheet, ok := stylesheets[filename]
			if !ok {
				continue
			}
			_, err = io.WriteString(w, stylesheet)
			if err != nil {
				return err
			}
		}

		return err
	})
}
