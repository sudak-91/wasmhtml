wasm: 
	@GOOS="js" GOARCH="wasm" go build -o ./example/selfhost/main.wasm ./example/selfhost/main.go