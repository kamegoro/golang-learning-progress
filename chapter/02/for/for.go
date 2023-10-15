package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	arr := [...]int{0, 1, 2, 3, 4, 5}

	for j := range arr {
		fmt.Println(j)
	}
}
