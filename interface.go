package main

import (
	"fmt"
)

func doTypeThings(i interface{}) {
	switch j := i.(type) {
	case nil:
		fmt.Println("nil", j)
	case int:
		fmt.Println("int", j)
	case bool:
		fmt.Println("bool", j)
	case string:
		fmt.Println("string", j)
	default:
		fmt.Println("nothing", j)
	}
}

func main() {
	var s *string
	fmt.Println(s == nil)
	var i interface{}
	fmt.Println(i == nil)
	i = s
	fmt.Println(i == nil)
	if _, ok := i.(*string); ok {
		fmt.Println("*string")
	}
	if _, ok := i.(int); ok {
		fmt.Println("int")
	}

	doTypeThings(1)
	doTypeThings(true)
	doTypeThings("haha")
}
