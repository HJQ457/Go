package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"Age"`
	Skill string `json:"Skill"` //序列化后可以输出小写
}

func main() {
	var a A
	var b B

	a = A(b) //可以转换，但是要求结构体的字段要完全一样（包括：名字，个数和类型）
	fmt.Println(a, b)

	//创建变量
	monster := Monster{"tom", 500, "aaa"}
	fmt.Println(monster)

	//将monser变量转化为json格式字符串
	jsonStr, err := json.Marshal(monster)

	if err != nil {
		fmt.Println("json 处理错误", err)
	}

	fmt.Println("jsonStr", string(jsonStr)) //json.Marshal方法默认返回的是[]byte，需要转换成string
}
