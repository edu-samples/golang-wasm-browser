package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("greet", js.FuncOf(greet))
	<-c
}

func greet(this js.Value, args []js.Value) interface{} {
	name := args[0].String()
	message := fmt.Sprintf("Hello, %s! Welcome to Go WebAssembly!", name)
	return js.ValueOf(message)
}
