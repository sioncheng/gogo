package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("compiler", runtime.Compiler)
	fmt.Println("arch", runtime.GOARCH)
	fmt.Println("go version", runtime.Version())
	fmt.Println("cpus", runtime.NumCPU())
	fmt.Println("goroutines", runtime.NumGoroutine())
}
