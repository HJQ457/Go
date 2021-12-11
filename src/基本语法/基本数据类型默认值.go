package main

import "fmt"

func main() {
	var a int     //0
	var b float32 //0
	var c float64 //0
	var d bool    //flase
	var e string  //""

	fmt.Println("a=", a, "b=", b, "c=", c, "d=", d, "e=", e)

	//%v是按照变量的值输出，原始输出
	fmt.Printf("a=%d,b=%f,c=%f,d=%v,e=%v", a, b, c, d, e)
}
