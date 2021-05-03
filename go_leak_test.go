package main

import "testing"

func TestLeak(t *testing.T) {
	leak()
}
