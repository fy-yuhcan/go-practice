package main

import "fmt"

//x := 1 //エラー

func main() {
	i, j := 1, 2
	c, python, java := true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

//varの代わりに:=を使って変数宣言を行うことができる
//関数外では:=を使うことはできない
//関数の外ではキーワードで始まる宣言（var,funcなど）が必要