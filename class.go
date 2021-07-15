package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func (c *Counter) DoNothing() {
	//
}

func main() {
	var c Counter
	fmt.Println(c)
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
	var cp *Counter = nil
	fmt.Println(cp)
	//will panic
	//cp.Increment()
	//fmt.Println(cp.String())
	//cp.DoNothing()
}
