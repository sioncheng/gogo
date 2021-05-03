package main

import (
	"fmt"
	"strconv"
)

const OneHand = 100

type Stock struct {
	Name  string
	Price float64
}

type StockPercent struct {
	Stock   Stock
	Percent float64
}

type StockPortfolio struct {
	Stock   Stock
	Shares  int64
	Money   float64
	Percent float64
}

func float2(f float64) (float64, error) {
	return strconv.ParseFloat(fmt.Sprintf("%.2f", f), 64)
}

func meetInvesmentExpection(total float64, spe []StockPercent) (float64, []StockPortfolio) {
	l := len(spe)
	result := make([]StockPortfolio, l)
	bb := make([]bool, l)
	restTotal := total
	percentTotal := 1.0

	for i := 0; i < l; i++ {
		result[i] = StockPortfolio{spe[i].Stock, 0, 0.0, 0.0}
		bb[i] = true
	}

	for {
		spent := 0.0
		b := false
		p := 0.0
		for i := 0; i < l; i++ {
			if !bb[i] {
				continue
			}
			pm := restTotal * (spe[i].Percent / percentTotal)
			hm1 := spe[i].Stock.Price * OneHand
			shares := int64(pm / hm1)
			if shares >= 1 {
				hm := hm1 * float64(shares)
				sp := result[i]
				sp.Shares = sp.Shares + shares*OneHand
				sp.Money = sp.Money + hm
				spent = spent + hm
				b = true
				result[i] = sp
				p = p + spe[i].Percent
			} else {
				bb[i] = false
			}
		}
		restTotal = restTotal - spent
		percentTotal = p

		if !b {
			break
		}
	}

	for {
		b := false
		for i := 0; i < l; i++ {
			pm := restTotal
			hm1 := spe[i].Stock.Price * OneHand
			shares := int64(pm / hm1)
			if shares >= 1 {
				hm := hm1 * float64(shares)
				sp := result[i]
				sp.Shares = sp.Shares + shares*OneHand
				sp.Money = sp.Money + hm
				result[i] = sp
				restTotal = restTotal - hm
				b = true
			}
		}

		if !b {
			break
		}
	}

	for i := 0; i < len(result); i++ {
		sp := result[i]
		sp.Percent, _ = float2(sp.Money / total)
		result[i] = sp
	}

	return restTotal, result
}

func main() {

	s1 := Stock{"s1", 34.0}
	s2 := Stock{"s2", 44.5}
	s3 := Stock{"s3", 23.4}
	s4 := Stock{"s4", 11.3}
	s5 := Stock{"s5", 50}

	sp1 := StockPercent{s1, 0.50}
	sp2 := StockPercent{s2, 0.10}
	sp3 := StockPercent{s3, 0.10}
	sp4 := StockPercent{s4, 0.15}
	sp5 := StockPercent{s5, 0.15}

	sps := []StockPercent{sp1, sp2, sp3, sp4, sp5}

	rest, sp := meetInvesmentExpection(20000, sps)

	fmt.Println("expection", sps)
	fmt.Println("portfolio", sp)
	fmt.Println("rest", rest)

	rest, sp = meetInvesmentExpection(30000, sps)

	fmt.Println("expection", sps)
	fmt.Println("portfolio", sp)
	fmt.Println("rest", rest)

	rest, sp = meetInvesmentExpection(40000, sps)

	fmt.Println("expection", sps)
	fmt.Println("portfolio", sp)
	fmt.Println("rest", rest)

	rest, sp = meetInvesmentExpection(50000, sps)

	fmt.Println("expection", sps)
	fmt.Println("portfolio", sp)
	fmt.Println("rest", rest)

	rest, sp = meetInvesmentExpection(60000, sps)

	fmt.Println("expection", sps)
	fmt.Println("portfolio", sp)
	fmt.Println("rest", rest)

}
