package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "D:\\gocode\\src\\基本语法\\file1\\test.txt"
	//ioutil.ReadFile一次性将文件读取到位
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Read File Err=%v", err)
	}

	//把读取到的内容输出到终端
	fmt.Println(content) //输出为[]byte
	fmt.Printf("%v", string(content))
}
