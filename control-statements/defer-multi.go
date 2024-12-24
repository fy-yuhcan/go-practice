package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i :=0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
//deferに渡した関数が複数ある時、呼び出しはスタックされる
//関数がreturnするときにdeferへ渡した関数はLIFOの順番で実行される
//LIFO = Last In First Out 後入れ先だし