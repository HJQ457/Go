package main

import "fmt"

func main() {
	//使用select可以解决管道读取数据的阻塞问题

	//定义一个管道，包含10个int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	//定义一个管道 5个string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//实际开发中不确定什么时间关闭管道，可以使用select方式解决
	var flag bool = true
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intchan读取的数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
		default:
			fmt.Printf("都取不到\n")
			flag = false
		}
		if !flag {
			break
		}
	}
}
