package main

import (
	"fmt"
	"time"
)

func writeData(intChar chan int) {
	for i := 0; i < 50; i++ {
		intChar <- i
		fmt.Printf("writeData写入数据 %v\n", i)
		//time.Sleep(time.Second)
	}
	close(intChar)
}

func readData(intChar chan int, exitChan chan bool) {
	for {
		x, ok := <-intChar
		if !ok {
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("readData读到数据 %v\n", x)
	}
	exitChan <- true
	close(exitChan)
}
func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	//time.Sleep(10 * time.Second)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
