package main

import "fmt"

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisments
)

func main() {
	p := Personal
	fmt.Println(p)
}
