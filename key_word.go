package main

import (
	"fmt"
	"sort"
)

type KeywordCount struct {
	Name  string
	Count int
}

var kMap map[string]int = make(map[string]int)
var kList []KeywordCount = make([]KeywordCount, 1)
var n int = 0
var c int = 1

func addKeywords(words string) {
	if val, ok := kMap[words]; ok {
		kc := kList[val]
		kList[val] = KeywordCount{Name: words, Count: kc.Count + 1}
	} else {
		kMap[words] = n
		kc := KeywordCount{Name: words, Count: 1}
		kList[n] = kc
		n = n + 1
		if n >= c {
			nc := c * 2
			nkList := make([]KeywordCount, nc)
			for i := 0; i < n; i++ {
				nkList[i] = kList[i]
			}
			kList = nkList
			c = nc
		}
	}
}

func getTop100() []KeywordCount {
	//
	sort.Slice(kList, func(i, j int) bool {
		return kList[i].Count > kList[j].Count
	})
	max := 100
	if n < max {
		max = n
	}
	return kList[0:max]
}

func main() {
	addKeywords("a")
	addKeywords("b")
	addKeywords("a")
	addKeywords("c")
	fmt.Println(getTop100())
}
