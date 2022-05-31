package main

import "fmt"

func putNum(intChar chan int) {
	for i := 1; i <= 8000; i++ {
		intChar <- i
	}
	close(intChar)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	//var num int
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag := true //假设是素数
		//判断是否为素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			//将数字放到rimeChan
			primeChan <- num
		}
	}
	fmt.Println("有一个primeNum协程因为读取不到数据，退出")
	exitChan <- true
}
func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000) //放结果
	//标识退出的管道
	exitChan := make(chan bool, 4) //4个

	//开启一个协程，向intChar放入1-8000
	go putNum(intChan)

	//开启一个协程，从intChar取出数据，并判断是否为素数，如果是，放入primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
}
