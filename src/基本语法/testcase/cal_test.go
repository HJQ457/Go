package main

import (
	"testing"
)

//测试用例文件名必须以_test.go结尾

//编写测试用例，去测试appUpper是否正确
func TestAddUpper(t *testing.T) {
	//调用
	res := addUpper(10)

	if res != 555 {
		//fmt.Println("addupper执行错误")
		t.Fatalf("addupper执行错误，期望值%v,实际值%v", 555, res) //输出日志，并停止程序
	}

	//如果正确，输出日志
	t.Log("addupper执行正确")
}
