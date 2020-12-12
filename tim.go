package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("epoch time", now.Unix())
	fmt.Println(now, " - ", now.Format(time.RFC3339))
}
