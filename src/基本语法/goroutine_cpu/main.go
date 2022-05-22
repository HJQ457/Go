package main

import (
	"fmt"
	"runtime"
)

func main() {

	//获取cpu个数
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)

	//设置使用的cpu个数
	runtime.GOMAXPROCS(cpuNum)

}
