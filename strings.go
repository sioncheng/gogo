package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.ToLower("HELLo"))
	fmt.Println(strings.Title("this is a title"))
	fmt.Println(strings.EqualFold("HELLO", "Hello"))
	fmt.Println(strings.HasPrefix("hello", "h"))
	fmt.Println(strings.Index("Hello", "l"))
	fmt.Println(strings.Count("Hello", "l"))
	fmt.Println(strings.Repeat("o", 5))
	fmt.Println(strings.TrimSpace("hello go "))
	fmt.Println(strings.TrimLeft("hello go ", " "))
	fmt.Println(strings.TrimRight("hello go ", " "))

	var s string = "hello there"
	var s2 string = s[4:7]
	fmt.Println(s)
	fmt.Println(s2)
	var a rune = 'x'
	var b byte = 'y'
	fmt.Println(string(a), string(b))
}
