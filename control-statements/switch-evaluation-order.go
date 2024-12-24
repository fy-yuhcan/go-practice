package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
//switch文は上から下に評価される
//case節は評価された後にマッチするものがあれば実行される
//case節には定数が必要である
//case節には重複する値があってはならない
//default節は必ず最後に書く必要がある
//switch文の条件は省略可能で、その場合trueとなる
