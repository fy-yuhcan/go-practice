package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}

// 構造体はフィールドの集まりであり、フィールドはフィールド名と型を持つ
// structsはフィールドを波括弧{}で囲んで表現する