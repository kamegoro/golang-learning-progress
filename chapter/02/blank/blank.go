package main

import "fmt"

func f() (int, int) {
	return 100, 200
}

func main() {
	// コロンイコールはvarを省略するため
	i, _ := f()
	fmt.Println(i)
}
