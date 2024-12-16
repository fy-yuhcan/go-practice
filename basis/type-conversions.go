package main

import (
	"fmt"
	"math"
)

func main() {
	var x,y int = 3,4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	var s string = fmt.Sprint(z)
	// var s string = string(z) //エラー
	fmt.Println(x,y,f,z,s)
}

//int(x)やfloat64(x)のように型変換を行うことができる
//ただし、string(z)はできず、ftm.Sprint(z)のようにする必要がある
//型変換は明示的に行う必要がある