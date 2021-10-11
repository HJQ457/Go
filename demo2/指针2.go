package main

import "fmt"

func main() {
	var a int = 300
	var b int = 400

	var ptr *int = &a
	*ptr = 100 //a = 100
	ptr = &b
	*ptr = 200 //b = 200

	fmt.Printf("a=%v,b=%v,*ptr=%v\n", a, b, *ptr)
}
