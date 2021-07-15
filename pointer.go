package main

import "fmt"

func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

func main() {
	x := 1
	pointerX := &x
	fmt.Println(x, pointerX, *pointerX)
	*pointerX = 2
	fmt.Println(x, pointerX, *pointerX)
	failedUpdate(&x)
	fmt.Println(x)
	update(&x)
	fmt.Println(x)
}
