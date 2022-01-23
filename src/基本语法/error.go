package main

import (
	"errors"
	"fmt"
)

func main() {
	//test4()
	//for i := 1;i<= 5; i++ {
	//	fmt.Println("main()下面的代码")
	//	time.Sleep(time.Second)	//此函数不能直接传入值
	//}

	//测试自定义错误的使用
	test5()
}

func test4() {
	//使用defer + recover来捕获和处理异常
	defer func() {
		err := recover() //recover是内置函数，可以捕获异常
		if err != nil {  //说明捕获到错误
			fmt.Println("error=", err)
		}
	}() //()为匿名函数的调用

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

//函数读取配置文件init.conf的信息，如果文件名不正确，抛出一个自定义错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("文件读取错误")
	}
}

func test5() {
	err := readConf("config.ini")
	if err != nil {
		//如果读取文件发生错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test5()继续执行")
}
