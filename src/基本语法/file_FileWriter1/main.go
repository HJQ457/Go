package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "D:\\gocode\\src\\基本语法\\file_FileWriter1\\test.txt"

	//打开文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err = %v", err)
		return
	}

	defer file.Close()

	//写入文件，用带缓存的writer
	str := "hello Gardon\n" // \n表示换行
	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存的，因此在调用writer方法时，内容是先写入缓存的，需要flush方法落到硬盘上
	writer.Flush()
}
