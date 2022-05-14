package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//将test1.txt的内容读取到内存
	file1 := "D:\\gocode\\src\\基本语法\\file_copy\\text1.txt"
	file2 := "D:\\gocode\\src\\基本语法\\file_copy\\test2.txt"

	data, err1 := ioutil.ReadFile(file1)
	if err1 != nil {
		fmt.Println("read file err = %v", err1)
	}

	err2 := ioutil.WriteFile(file2, data, 0666)
	if err2 != nil {
		fmt.Println("write file err = %v", err2)
	}
}
