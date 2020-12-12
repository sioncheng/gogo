package main

import "fmt"

func main() {
	s := make([]int, 2)
	fmt.Println("cap and len", cap(s), len(s))
	s[0] = 1
	s[1] = 2
	fmt.Println("cap and len", cap(s), len(s))
	s = append(s, 3)
	fmt.Println("cap and len", cap(s), len(s))
}
