package main

import "fmt"

func main() {
	sum :=1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
//セミコロンを省略することができる、while文のように使うことができる

//無限ループ
func main() {
	for {
	}
}
//条件を省略すると無限ループになる