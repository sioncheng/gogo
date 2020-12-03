package main

import (
	"errors"
	"log"
)

func isError(a, b int) error {
	if a == b {
		return errors.New("Error happened")
	} else {
		return nil
	}
}

func main() {
	err := isError(1, 1)
	if nil != err {
		log.Panic(err)
	}
}
