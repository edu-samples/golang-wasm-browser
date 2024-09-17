
This will start a local web server on port 8080. Open your web browser and navigate to `http://localhost:8080` to see the example in action.

## Project Structure

- `main.go`: The Go code that will be compiled to WebAssembly.
- `index.html`: The HTML file that loads and uses the WebAssembly module.
- `wasm_exec.js`: JavaScript glue code provided by Go (copied from Go installation).
- `Makefile`: Automates the build process.

## How it Works

1. The Go code in `main.go` defines a `greet` function that takes a name as input and returns a greeting message.
2. The code is compiled to WebAssembly using the `GOOS=js GOARCH=wasm go build` command.
3. The HTML file loads the `wasm_exec.js` script and the compiled WebAssembly module.
4. When the user enters a name and clicks the "Greet" button, the JavaScript code calls the `greet` function from the WebAssembly module and displays the result.

## Cleaning Up

To remove the generated files, run:

