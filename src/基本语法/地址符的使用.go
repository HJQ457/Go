package main

import "fmt"

func main() {
	a := 100
	fmt.Println("a的地址为：", &a)

	var ptr *int = &a
	fmt.Println("ptr指向的值为：", *ptr)
}
