package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	monster := Monster{
		Name:     "Tom",
		Age:      100,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "Tom skill",
	}

	//将moster序列号
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列号失败：%v", err)
	}

	//输出序列号后结果
	fmt.Printf("%v", string(data))
}

//将map序列号
func testMap() {
	var a map[string]interface{}

	//使用map前需要make
	a = make(map[string]interface{})
	a["name"] = "Sam"
	a["age"] = "18"
	a["address"] = "北京"

	//将map进行序列号
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化失败：%v", err)
	}

	//输出序列号后结果
	fmt.Printf("%v", string(data))
}

//切片序列号
func testSlice() {
	var slice []map[string]interface{}

	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "Jack"
	m1["age"] = 10
	m1["address"] = "上海"

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "Marry"
	m2["age"] = 20
	m2["address"] = "北京"

	slice = append(slice, m1)
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列号失败：%v", err)
	}

	//输出序列号后结果
	fmt.Printf("%v", string(data))
}

func main() {
	testStruct()
	fmt.Println()
	testMap()
	fmt.Println()
	testSlice()
}
