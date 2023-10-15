package main

import "fmt"

func main() {
	holiday(1, "元旦", "成人の日")
	holiday(2, "建国記念の日")
	holiday(3, "春分の日")
}

// 論理値(bool)	%t
// 符号付き整数(int, int8など)	%d
// 符号なし整数(uint, uint8など)	%d
// 浮動小数点数(float64など)	%g
// 複素数(complex128など)	%g
// 文字列(string)	%s
// チャネル(chan)	%p
// デフォルト形式	%v
// %そのものを出力	%%
// 文字		%c
// 型表示	%T
// 2進数	%b
// ポインタ(pointer)	%p

func holiday(month int, days ...string) {
	fmt.Printf("%d月の祝日は%d日あります。\n", month, len(days))

	for _, day := range days {
		fmt.Println(day)
	}

	fmt.Println()
}
