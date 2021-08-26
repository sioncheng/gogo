package algr

import "testing"

func TestFibonacci(t *testing.T) {
	r0 := fibonacci(0)
	if 0 == r0 {
		t.Log("ok")
	} else {
		t.Error("not 0")
	}
	r1 := fibonacci(1)
	if 1 == r1 {
		t.Log("ok")
	} else {
		t.Error("not 1")
	}
	r2 := fibonacci(2)
	if 1 == r2 {
		t.Log("ok")
	} else {
		t.Error("not 1")
	}
}
