package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	parent, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	child, childCancel := context.WithTimeout(parent, 3*time.Second)
	defer childCancel()
	start := time.Now()
	<-child.Done()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
