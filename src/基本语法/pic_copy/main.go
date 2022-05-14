package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcfile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file err = %v", err)
	}

	defer srcfile.Close()

	reader := bufio.NewReader(srcfile)

	dstfile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}

	defer dstfile.Close()

	writer := bufio.NewWriter(dstfile)

	return io.Copy(writer, reader)
}

func main() {

	srcfile := "D:\\gocode\\src\\基本语法\\pic_copy\\test1.txt"
	dstfile := "D:\\gocode\\src\\基本语法\\pic_copy\\test2.txt"

	_, err := CopyFile(dstfile, srcfile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Println("copy err = %v", err)
	}

}
