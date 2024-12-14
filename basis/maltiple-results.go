package main

import "fmt"

func swap(x, y,z string) (string, string,string) {
	return z,y, x
}

func main() {
	a, b,c := swap("hello", "world","!")
	fmt.Println(a, b,c)
}

//private関数なので先頭は小文字で”swap”となる
//mainパッケージのmain関数が実行される
//swap関数は3つの引数を受け取り、3つの戻り値を返す
//swap関数は3つの引数を逆の順番で返す
//main関数でswap関数を呼び出し、3つの戻り値を受け取り、それぞれの変数に代入する
