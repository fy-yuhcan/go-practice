package main

import (
	"fmt"
	"strconv"
)

func add(x,y int,z string)int {
	i,_ := strconv.Atoi(z)
	return x + y + i
}

func main() {
	fmt.Println(add(42, 13, "1"))
}

//private関数なので先頭は小文字で”add”となる
//mainパッケージのmain関数が実行される
//引数には型を指定する
//戻り値にも型を指定する
//strconvパッケージをインポートすることで、文字列を数値に変換することができる
//strconv.Atoi()関数は文字列を整数に変換する
//_:はブランク識別子で、不要な値を無視するために使用される
