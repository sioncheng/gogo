package main

import (
	"fmt"
)

func main() {
	var nilMap map[string]int

	fmt.Println(len(nilMap))
	fmt.Println(nilMap["k"])

	teams := map[string][]string{
		"Orcas": []string{"Fred", "Ralph", "Bijou"},
		"Lions": []string{"Sarah", "Peter", "Billie"},
	}

	fmt.Println(teams)

	ages := make(map[int][]string, 10)
	ages[20] = []string{"Waldo", "Raul", "Ze"}
	fmt.Println(len(ages))

	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"], totalWins["Lions"])
	totalWins["Kittens"]++
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])

	intSet := map[int]bool{}
	vals := []int{1, 4, 5, 3, 2}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	if intSet[4] {
		fmt.Println("4 is in the set")
	}
}
