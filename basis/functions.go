package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

//private関数なので先頭は小文字で”add”となる
//mainパッケージのmain関数が実行される　
//引数には型を指定する
//戻り値にも型を指定する