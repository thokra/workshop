package main

import (
	"runtime"

	"fmt"
)

func main() {
	fmt.Println(runtime.Version())
	fmt.Print("Hello, World!")
}
