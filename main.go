package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("greet", js.FuncOf(greet))
	js.Global().Set("fibonacci", js.FuncOf(fibonacci))
	<-c
}

func greet(this js.Value, args []js.Value) interface{} {
	name := args[0].String()
	message := fmt.Sprintf("Hello, %s! Welcome to Go WebAssembly!", name)
	return js.ValueOf(message)
}

func fibonacci(this js.Value, args []js.Value) interface{} {
	n, err := strconv.Atoi(args[0].String())
	if err != nil || n < 0 {
		return js.ValueOf("Please enter a valid non-negative integer")
	}
	result := fib(n)
	return js.ValueOf(fmt.Sprintf("Fibonacci(%d) = %d", n, result))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
