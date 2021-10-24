package main

import "fmt"

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool

	//方式一
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	fmt.Println("请输入薪水：")
	fmt.Scanln(&sal)

	fmt.Println("请输入是否通过：")
	fmt.Scanln(&isPass)

	fmt.Printf("name=%v,age=%v,sal=%v,isPass=%v\n", name, age, sal, isPass)

	//方式二
	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试，请用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("name=%v,age=%v,sal=%v,isPass=%v", name, age, sal, isPass)
}
