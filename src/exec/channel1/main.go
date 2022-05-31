package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) //关闭管道

	//intChan <- 300	此处会报错，因为管道已经关闭，不能写入数据

	n1 := <-intChan
	fmt.Println(n1)

	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	//取出管道中的值，不能使用for循环，必须使用for range
	for v := range intChan2 {
		fmt.Println(v)
	}
}
