package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
//goの基本型は以下の通り
//bool
//string
//int int8 int16 int32 int64　符号あり整数
//uint uint8 uint16 uint32 uint64 uintptr　符号なし整数
//byte // uint8 の別名
//rune // int32 の別名
//float32 float64　浮動小数点数
//complex64 complex128　複素数
//特別な理由がなければintを使う