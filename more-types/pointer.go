package main

import "fmt"

func main() {
	i,j := 42,2701

	p := &i //ポインタpはiのアドレスを指す
	fmt.Println(*p) //ポインタpを通してiの値を読み出す
	*p = 21 //ポインタpを通してiに21を代入する
	fmt.Println(i) //iの値が変わっている

	p = &j //ポインタpはjのアドレスを指す
	*p = *p / 37 //ポインタpを通してjの値を読み出し、それを37で割る
	fmt.Println(j) //jの値が変わっている
}
