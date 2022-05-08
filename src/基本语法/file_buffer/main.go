package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("D:/gocode/src/基本语法/file1/test.txt")
	if err != nil {
		fmt.Println("open file err = ", err)
	}

	defer file.Close() //函数退出时关闭执行defer，关闭流

	//创建一个reader,默认缓冲区4096
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到换行符就结束

		if err == io.EOF { //io.EOF表示文件的末尾
			break
		}

		//输出内容
		fmt.Print(str)
	}

	fmt.Println("文件读取结束")
}
