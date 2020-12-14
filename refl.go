package main

import (
	"fmt"
	"reflect"
)

type xs struct {
	A float64
	B byte
	C int
}

func main() {
	x := 100
	xValue := reflect.ValueOf(&x)
	xType := reflect.TypeOf(&x)

	fmt.Println(xValue, xType)
	fmt.Println(xValue.Type(), xValue.Elem())

	xs1 := xs{A: 6.0, B: 1, C: 2}
	xsElem1 := reflect.ValueOf(&xs1).Elem()
	fmt.Println(xsElem1)

	fmt.Printf(" type: %s\n", xsElem1.Type())
	fmt.Printf(" num fields: %d\n", xsElem1.NumField())

	xs1Type := reflect.TypeOf(&xs1).Elem()
	for i := 0; i < xsElem1.NumField(); i++ {
		f := xsElem1.Field(i)
		fmt.Println(f.String())
		fmt.Println(f.Type())
		fmt.Println(f.Interface())
		fmt.Println(xs1Type.Field(i).Name)
		fmt.Println(" ---- ")
	}

}
