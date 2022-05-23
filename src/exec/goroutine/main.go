package main

import (
	"fmt"
	"sync"
	"time"
)

//计算1-200各个数阶乘，放到map中

var (
	myMap map[int]int = make(map[int]int, 10)
	lock  sync.Mutex
)

//lock是全局互斥锁
//lock sync.Mutex
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res += i
	}

	//将res放到map中
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
func main() {
	//开启多个协程
	for i := 0; i <= 100; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 2)

	//输出结果
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}
