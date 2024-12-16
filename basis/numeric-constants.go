package main

import "fmt"

const (
	big  = 1 << 100 // 2の100乗
	small = big >> 99 // 2
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(small))
	fmt.Println(needFloat(small))
	fmt.Println(needFloat(big))
}
//定数は未定型であり必要におうじて適切な型に変換される