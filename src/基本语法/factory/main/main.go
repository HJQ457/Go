package main

import (
	"fmt"
	"基本语法/factory/model"
)

func main() {
	var stu = model.Newstudent("tom", 65.5)
	fmt.Println(*stu)
	fmt.Println((*stu).Name)
	fmt.Println(stu.Getscore())
}
