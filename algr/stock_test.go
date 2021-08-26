package algr

import "testing"

func TestStock1(t *testing.T) {
	prices := []int{1, 2, 3, 9, 0, 10}
	s := stock1(prices)
	if s != 10 {
		t.Error("not 10")
	}

	prices = []int{1, 1, 2, 3, 12, 9, 0, 10}
	s = stock1(prices)
	if s != 11 {
		t.Error("not 11")
	}

	prices = []int{7, 6, 5, 4, 3, 2, 1}
	s = stock1(prices)
	if s != 0 {
		t.Error("not 0")
	}
}

func TestStockN(t *testing.T) {
	prices := []int{1, 10, 1, 10}
	s := stockn(prices)
	if s != 18 {
		t.Error("not 18")
	}
}
