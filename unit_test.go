package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("set up stuff for tests here")
	testTime = time.Now()
	exit := m.Run()
	fmt.Println("Clean up stuff after tests here")
	os.Exit(exit)
}

func TestFunc(t *testing.T) {
	fmt.Println("TestFunc uses stuff set up in TestMain", testTime)
}
