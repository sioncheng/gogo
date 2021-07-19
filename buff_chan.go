package main

import (
	"fmt"
	"time"
)

const concurrent = 10

func process(i int) int {
	time.Sleep(time.Duration(i * 100 * int(time.Millisecond)))
	return i * i
}

func main() {
	result := make(chan int, concurrent)

	for i := 0; i < concurrent; i++ {
		go func(v int) {
			result <- process(v)
		}(i)
	}

	for i := 0; i < concurrent; i++ {
		fmt.Println(<-result)
	}
}
