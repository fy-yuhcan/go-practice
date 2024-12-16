package main

import "fmt"

func main() {
	sum := 0
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
//for文の初期化ステートメントと後処理ステートメントは省略可能