
start_time=$(date +%s%3N)

bun run --bun ./build/render/rendering/main.ts -i ./project/src/lib/ -o ./compile/
./build/render_to_templ/main -in ./compile/ -out ./project/gen/

cd ./project
TEMPL_EXPERIMENT=rawgo templ generate ./
go build -tags "sonic avx" -ldflags="-s -w" -o ../ ./src/main.go

end_time=$(date +%s%3N)

echo "Build time: $((end_time - start_time)) seconds"
