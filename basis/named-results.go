package main

import "fmt"

func divide(num,denominator int) (quotient, remainder int) {
	quotient = num / denominator
	remainder = num % denominator
	return
}

func main() {
	fmt.Println(divide(42, 13))
}

//戻り値となる変数に名前をつけることができる、そしてreturnに何も書かなくても戻り値を返すことができる
//戻り値の名前を明示することでドキュメントとして使える
//naked returnは短い関数でのみ利用するべきである
