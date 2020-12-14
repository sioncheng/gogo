package intf

type Square struct {
	X float64
}

type Circle struct {
	R float64
}

func (s Square) Area() float64 {
	return s.X * s.X
}

func (c Circle) Area() float64 {
	return 3.14 * c.R * c.R
}
