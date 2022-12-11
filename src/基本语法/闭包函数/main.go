package main

import "fmt"

//使用闭包函数优化study
//每个人有不同的进度，将这个进度保存在【各自的闭包】中

func GetStudy(name string) func(int) int {

	var progress int = 0

	study := func(hours int) int {
		fmt.Printf("%s学习了%d小时\n", name, hours)
		progress += hours

		return progress
	}
	//fmt.Printf("Study的类型是%T", study)

	return study
}

func main() {
	studyFunc := GetStudy("小明")
	studyFunc(5)
	studyFunc(3)
	xiaomingProgress := studyFunc(2)
	fmt.Printf("小明学习进度是%d/10000\n", xiaomingProgress)
}
