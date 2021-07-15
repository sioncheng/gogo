package main

import "fmt"

func main() {
	x := 10
	if 1 == 1 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}
