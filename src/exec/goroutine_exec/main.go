package main

import "fmt"

func main() {

	//声明为只写
	var chan1 chan<- int
	chan1 = make(chan int, 3)
	chan1 <- 20

	//声明为只读
	var chan2 <-chan int
	chan2 = make(chan int, 3)
	fmt.Println(chan2)
}
