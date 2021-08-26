package stock

import (
	"testing"
)

func TestSimpleMaxProfit(t *testing.T) {
	prices := []int{1, 3, 0, 5, 7, 9, 1}
	s := StockSolution{}
	m := s.simpleMaxProfile(prices)
	if m == 9 {
		t.Log("ok")
	} else {
		t.Error("not 9")
	}
}

func TestAdvanceMaxProfit(t *testing.T) {
	prices := []int{1, 3, 0, 5, 7, 9, 1}
	s := StockSolution{}
	m := s.advanceMaxProfile(prices)
	if m == 9 {
		t.Log("ok")
	} else {
		t.Error("not 9")
	}
}

func TestSimpleMaxProfit1(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5, 6}
	s := StockSolution{}
	m := s.simpleMaxProfile(prices)
	if m == 5 {
		t.Log("ok")
	} else {
		t.Error("not 5")
	}
}

func TestAdvanceMaxProfit2(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5, 6}
	s := StockSolution{}
	m := s.advanceMaxProfile(prices)
	if m == 5 {
		t.Log("ok")
	} else {
		t.Error("not 5")
	}
}
