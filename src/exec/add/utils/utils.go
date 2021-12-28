package utils

import (
	"fmt"
	"strings"
)

func MakeSuffix(suffix string) func(string) string {
	return func(name string) string {
		//如果 name 没有指定后缀，则加上，否则返回原来的名字
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}

func AddUper() func(int) int {
	var n int = 10
	var str string = "hello"
	return func(x int) int {
		n = n + x
		str += "a"
		fmt.Println("str = ", str)
		return n
	}
}
