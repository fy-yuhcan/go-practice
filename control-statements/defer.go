package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
//deferは関数の実行を遅延させる
//returnするまで遅延される