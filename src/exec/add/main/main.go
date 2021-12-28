package main

import (
	"exec/add/utils"
	"fmt"
)

func main() {
	f := utils.AddUper()
	fmt.Println(f(1))
	fmt.Println(f(2))

	f1 := utils.MakeSuffix(".jpg")
	fmt.Println("文件处理后=", f1("winter"))
	fmt.Println("文件处理后=", f1("bird.avi"))
}
