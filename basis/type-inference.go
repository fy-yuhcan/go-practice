package main

import "fmt"

func main() {
	i:= 42
	f := 42.124
	b := true
	i2 := i
	fmt.Printf("type of i is %T\n", i)
	fmt.Printf("type of f is %T\n", f)
	fmt.Printf("type of b is %T\n", b)
	fmt.Printf("type of i2 is %T\n", i2)
}
//型推論してくれる
//型を省略して変数宣言を行うことができる


