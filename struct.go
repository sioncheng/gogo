package main

import "fmt"

type XYZ struct {
	X int
	Y int
	Z int
}

func NewXYZ(x, y, z int) *XYZ {
	return &XYZ{x, y, z}
}

func main() {

	xyz := XYZ{X: 1, Y: 2, Z: 3}
	fmt.Println(xyz)
	fmt.Println(xyz.X, xyz.Y, xyz.Z)

	xyz2 := NewXYZ(1, 2, 3)
	fmt.Println(xyz2)
	fmt.Println((*xyz2).X, (*xyz2).Y, (*xyz2).Z)

	xyz3 := new(XYZ)
	fmt.Println(xyz3)
	fmt.Println((*xyz3).X, (*xyz3).Y, (*xyz3).Z)

}
