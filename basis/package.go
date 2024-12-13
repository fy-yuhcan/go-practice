package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))
}

//mainパッケージのmain関数が実行される
//fmtパッケージとmath/randパッケージがインポートされる
//パッケージ名はimportパスの最後の要素で指定される
//math/randパッケージはmathパッケージのサブパッケージである
//デフォルトだとrand.seed(1)が設定されているため、rand.Intn(10)は常に同じ値を返す
//timeパッケージは現在時刻を取得するために使用される