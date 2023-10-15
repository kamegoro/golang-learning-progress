package main

import "fmt"

func main() {
	// int型変数のアドレスを格納できるポインタの型
	var ptr *int

	var i int = 12345

	ptr = &i

	fmt.Println("iのアドレス:", &i)
	// アドレスが入っているので*付けて出力しないと値が出ない
	fmt.Println("ptrの値（変数iのアドレス）:", ptr)

	fmt.Println("iの値:", i)
	fmt.Println("ポインタ経由のiの値:", *ptr)

	*ptr = 999

	fmt.Println("ポインタ経由で変更したiの値:", i)
}
