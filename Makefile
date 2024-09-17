.PHONY: all clean serve deploy

GOROOT := $(shell go env GOROOT)

all: main.wasm wasm_exec.js

main.wasm: main.go
	GOOS=js GOARCH=wasm go build -o main.wasm main.go

wasm_exec.js:
	cp "$(GOROOT)/misc/wasm/wasm_exec.js" .

clean:
	rm -f main.wasm wasm_exec.js

serve:
	python3 -m http.server 4545

deploy: all
	@echo "Project is ready for deployment"

.PHONY: install-goexec
install-goexec:
	go get -u github.com/shurcooL/goexec
