package algr

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func stock1(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	b := -1 * prices[0]
	s := 0
	bj := 0
	sj := 0
	for i := 1; i < n; i++ {
		cs := prices[i] + b
		if cs > s {
			s = cs
			sj = i
		}
		cb := -1 * prices[i]
		if b < cb {
			b = cb
			bj = i
		}
	}

	fmt.Printf("bj %d sj %d", bj, sj)
	fmt.Println()

	return s
}

func stockn(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	b := -1 * prices[0]
	s := 0
	for i := 1; i < n; i++ {
		cs := prices[i] + b
		if cs > s {
			s = cs
		}
		cb := s - prices[i]
		if b < cb {
			b = cb
		}
	}

	return s
}

func stock2(prices []int) (int, int) {
	n := len(prices)
	if n <= 1 {
		return 0, 0
	}

	b1 := -1 * prices[0]
	s1 := 0
	b2 := 0
	s2 := 0

	for i := 1; i < n; i++ {
		b1 = max(b1, -1*prices[i])
		s1 = max(s1, b1+prices[i])
		b2 = max(b2, s1+(-1*prices[i]))
		s2 = max(s2, prices[i]+b2)
	}

	return s1, s2
}
