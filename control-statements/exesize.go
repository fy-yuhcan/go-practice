package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 100; i++ {
		//ニュートン法で平方根を求める
		newZ := x - (x*x - z) / (2*x)
		fmt.Println(i, ":", newZ)
		//収束した時にループを抜ける
		if math.Abs(newZ-z) < 0.00000000001 {
            break
        }
		z = newZ
	}
	return z
}

func main() {
	fmt.Println("result:", Sqrt(2))
    fmt.Println("result:", Sqrt(3))
    fmt.Println("result:", Sqrt(4))
}