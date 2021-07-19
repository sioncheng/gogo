package main

import (
	"math/rand"
	"testing"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func SomeFuc() {
	time.Sleep(time.Duration(rnd.Intn(200)) * time.Millisecond)
}

func BenchmarkSomeFunc(b *testing.B) {

	SomeFuc()

}
