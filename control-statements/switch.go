package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default:
			// freebsd, openbsd,
			// plan9, windows...
			fmt.Printf("%s.\n", os)
	}
}

//switch文は、条件分岐を行うための制御構文
//switch文の条件式には、if文と同様に、比較演算子や論理演算子を使うことができる
//switch文の条件式には、if文と異なり、変数を指定することができる
//breakステートメントは自動的に挿入される

