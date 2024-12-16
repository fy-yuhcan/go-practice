package main

import "fmt"

func main() {	
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %v\n", i, f, b, s)
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

//初期化するとゼロ値が与えられる
//int型は0
//float型は0.0
//bool型はfalse
//string型は""（空文字列）