package stock

type StockSolution struct{}

func (s *StockSolution) simpleMaxProfile(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			p := prices[j] - prices[i]
			if p > max {
				max = p
			}
		}
	}
	return max
}

func (s *StockSolution) advanceMaxProfile(prices []int) int {
	max := 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		p := prices[i] - min
		if p > max {
			max = p
		}
	}

	return max
}

func (s *StockSolution) multiShotProfile(prices []int) int {
	mp := 0
	buy := -1 * prices[0]

	for i := 1; i < len(prices); i++ {
		p := mp - prices[i]
		if p > buy {
			buy = p
		}
	}
}
