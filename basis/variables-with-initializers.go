package main

import "fmt"

var i,j int = 1,2

func main() {
	var c,python,java = true,false,"no!"
	fmt.Println(i,j,c,python,java)
}

//変数宣言はvarで行う
//変数名の後に型を指定する
//複数の変数を同時に宣言することも可能
//変数宣言時に初期化子を指定することも可能
//初期化子が指定されている場合、型を省略することができる
//関数内での変数宣言は省略形で行うことができる