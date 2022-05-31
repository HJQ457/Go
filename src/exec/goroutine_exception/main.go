package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world")
	}
}

func test() {

	//匿名函数捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test()发生错误", err)
		}

	}()

	var myMap map[int]string
	myMap[0] = "jack"
}
func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok = ", i)
		time.Sleep(time.Second * 1)
	}
}
