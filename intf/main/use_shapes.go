package main

import (
	"fmt"

	"github.com/sioncheng/gogo/intf"
)

func tellShape(s intf.Shape) {
	switch v := s.(type) {
	case intf.Square:
		fmt.Println("square", v)
	case intf.Circle:
		fmt.Println("circle", v)
	default:
		fmt.Println("else")
	}
}

func main() {

	s := intf.Square{X: 100}
	c := intf.Circle{R: 100}

	var si intf.Shape = s
	var sc intf.Shape = c

	fmt.Println(si.Area())
	fmt.Println(sc.Area())

	tellShape(si)
	tellShape(sc)

}
