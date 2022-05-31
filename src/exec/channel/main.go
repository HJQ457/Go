package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main() {
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom jack"

	cat := Cat{"huamao", 10}

	allChan <- cat

	//希望获得管道中第三个元素，需要将前两个推出
	<-allChan
	<-allChan

	newCat := <-allChan
	fmt.Printf("类型为%T,值为%v\n", newCat, newCat)

	//类型断言
	a := newCat.(Cat)
	fmt.Println(a.Name)
}
