package main

import (
	bufio "bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	filepath := "D:\\gocode\\src\\基本语法\\file_FileWriter2\\text.txt"

	//覆盖写
	file1, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 06666)
	if err != nil {
		fmt.Println("open file err = %v", err)
	}

	defer file1.Close()

	str1 := "hello hjq\n"

	writer1 := bufio.NewWriter(file1)

	for i := 0; i < 5; i++ {
		writer1.WriteString(str1)
	}

	writer1.Flush()

	//追加写
	file2, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("open file err = %v", err)
	}

	defer file2.Close()

	//先读取文件的内容，显示在终端
	reader2 := bufio.NewReader(file2)

	for {
		readString, err := reader2.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(readString)
	}

	writer2 := bufio.NewWriter(file2)

	str2 := "追加"

	writer2.WriteString(str2)

	writer2.Flush()
}
