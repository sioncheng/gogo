package main

import (
	"fmt"
	"time"
)

type Searcher func(string) string

func searchData(s string, searchers []Searcher) (string, string) {
	done := make(chan struct{})
	result := make(chan string, 2)

	for _, searcher := range searchers {
		go func(ss Searcher) {
			select {
			case result <- ss(s):
			case <-done:
			}
		}(searcher)
	}

	r1 := <-result
	r2 := <-result
	close(done)
	return r1, r2
}

func createSearcher(i int) Searcher {
	return func(s string) string {
		time.Sleep(time.Duration(i * 100))
		return fmt.Sprintf("%d %s", i*100, s)
	}
}

func main() {
	var searchers []Searcher
	for i := 0; i < 5; i++ {
		searchers = append(searchers, createSearcher(i))
	}

	fmt.Println(searchData("hello", searchers))
}
