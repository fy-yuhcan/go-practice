//可視性　つまりアクセス制御
package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(math.pi) //エラー
	fmt.Println(math.Pi)
}

//math.piはエラーになる,
//goでは可視性が先頭文字の大文字か小文字かで区別される
//math.Piは大文字で始まるため、外部パッケージからアクセス可能で、math.piは小文字で始まるため、外部パッケージからアクセス不可能
//main->mathパッケージの時は大文字にする必要がある