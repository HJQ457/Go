package main

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		fmt.Println("存在")
		return true, nil
	}

	if os.IsNotExist(err) {
		fmt.Println("不存在")
		return false, nil
	}
	return false, nil
}

func main() {
	PathExists("D:\\gocode\\src\\基本语法\\file_copy\\text1.txt")
}
