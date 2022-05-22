package main

import (
	"fmt"
	"strconv"
	"time"
)

//编写一个函数，输出hello world
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
func main() {

	go test() //开启了一个协程

	for i := 0; i < 10; i++ {
		fmt.Println("main() hello world " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
