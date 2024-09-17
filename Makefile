.PHONY: all clean serve

GOROOT := $(shell go env GOROOT)

all: main.wasm wasm_exec.js

main.wasm: main.go
	GOOS=js GOARCH=wasm go build -o main.wasm main.go

wasm_exec.js:
	cp "$(GOROOT)/misc/wasm/wasm_exec.js" .

clean:
	rm -f main.wasm wasm_exec.js

serve:
	go run -v github.com/shurcooL/goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'

.PHONY: install-goexec
install-goexec:
	go get -u github.com/shurcooL/goexec
