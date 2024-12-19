package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

// ifステートメントは、条件式の前に短いステートメントを書くことができる
// このステートメントで宣言された変数は、ifスコープ内でのみ有効