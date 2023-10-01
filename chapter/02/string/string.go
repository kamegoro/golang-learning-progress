package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var en string = "golang"
	var ja string = "GO言語"

	// 漢字はbyte数が異なる
	fmt.Println(en, " len", len(en))
	fmt.Println(ja, " len", len(ja))

	fmt.Println(ja, " len", utf8.RuneCountInString(ja))
}
