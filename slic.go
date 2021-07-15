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

	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("z", z)
}
