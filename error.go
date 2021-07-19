package main

import (
	"errors"
	"fmt"
)

func main() {
	err0 := errors.New("error0")
	err1 := fmt.Errorf("error :%w", err0)
	fmt.Println(err0)
	fmt.Println(err1)
	if wrappedErr := errors.Unwrap(err1); wrappedErr != nil {
		fmt.Println("wrappedErr", wrappedErr)
	}
	errors.As(A, a)
}
