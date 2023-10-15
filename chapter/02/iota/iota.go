package main

import "fmt"

const (
	one   = iota
	two   = iota
	three = iota
	four  = iota
)

func main() {
	fmt.Println(one, two, three, four)
}
