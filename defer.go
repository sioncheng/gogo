package main

import "fmt"

func returnValues() int {
	var result int = 0
	defer func() {
		result++
	}()
	return result
}

func namedReturnVaues() (result int) {
	result = 0
	defer func() {
		result++
	}()
	return result
}

func returnSlice() []int {
	result := make([]int, 5, 5)
	defer func() {
		result[0] = 1
	}()
	return result
}

func main() {
	fmt.Println(returnValues())
	fmt.Println(namedReturnVaues())
	fmt.Println(returnSlice())
}
