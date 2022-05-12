package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("D:/gocode/src/基本语法/file1/test.txt")
	if err != nil {
		fmt.Println("open file err = ", err)
	}

	//输出文件
	fmt.Println(file)

	err = file.Close()

	if err != nil {
		fmt.Println("close file err = ", err)
	}
}
