package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Pat", 37},
		{"Tracy", 43},
		{"Fred", 30},
	}
	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)
}
