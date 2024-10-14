go install github.com/a-h/templ/cmd/templ@latest

bun install
go -C ./build/render_to_templ/ build -ldflags="-s -w" -o ./main.exe ./cmd/main.go
go -C ./build/builder/ build -ldflags="-s -w" -o ./build.exe ./build.go
./build/builder/build.exe