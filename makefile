wasm: 
	@export WASM_DEBUG=TRUE;\
	echo $$WASM_DEBUG;\
	GOOS="js" GOARCH="wasm" go build -o ./example/selfhost/main.wasm ./example/selfhost/main.go