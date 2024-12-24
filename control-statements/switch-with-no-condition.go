package main

import(
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon.")
		default:
			fmt.Println("Good evening.")
	}
}

//条件のないswitch文は、if-then-elseのように使うことができる