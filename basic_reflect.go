package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int    `mytag:"A"`
	B string `mytag:"B"`
}

func main() {
	fmt.Println("//// type ")
	var x int
	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name())
	f := Foo{1, "1"}
	ft := reflect.TypeOf(f)
	fmt.Println(ft.Name())
	for i := 0; i < ft.NumField(); i++ {
		field := ft.Field(i)
		fmt.Println(field.Name, field.Type.Name(), field.Tag.Get("mytag"))
	}
	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.Name())
	fmt.Println(xpt.Kind())
	fmt.Println(xpt.Elem().Name())
	fmt.Println(xpt.Elem().Kind())

	x2 := reflect.New(reflect.TypeOf(x))
	x2.Elem().SetInt(20)
	fmt.Println(x2)

	fmt.Println("")
	fmt.Println("//// value")
	s := []string{"a", "b", "c"}
	sv := reflect.ValueOf(s)
	fmt.Println(sv.Kind())
	s2 := sv.Interface().([]string)
	fmt.Println(s2)

	i := 10
	iv := reflect.ValueOf(&i)
	ivv := iv.Elem()
	ivv.SetInt(20)
	fmt.Println(i)

	fmt.Println("//// make")
	stringType := reflect.TypeOf("a")
	stringSliceType := reflect.TypeOf([]string(nil))
	ssv := reflect.MakeSlice(stringSliceType, 0, 10)
	sv = reflect.New(stringType).Elem()
	sv.SetString("hello")
	ssv = reflect.Append(ssv, sv)
	ss := ssv.Interface().([]string)
	fmt.Println(ss)
}
