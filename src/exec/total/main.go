package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	fileName := "D:\\gocode\\src\\exec\\total\\test.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err = %v", err)
	}

	defer file.Close()

	//创建实例
	var count CharCount

	//创建reader
	reader := bufio.NewReader(file)

	//读取文件内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		//遍历str
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				count.ChCount++
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' && v == '\t':
				count.NumCount++
			case v >= '0' && v <= '9':
				count.OtherCount++
			}
		}
	}
	fmt.Printf("字符个数为%v，数字个数为%v，空格个数为%v，其他字符个数%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
