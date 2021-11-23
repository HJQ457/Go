package main

import "fmt"

func main() {
	i := 1
	for ; i <= 10; i++ {
		fmt.Println("hello", i)
	}

	//for的第二种写法
	var j int = 1
	for j <= 10 {
		fmt.Println("hi", j)
		j++
	}

	//for循环的第三种写法，通常配合break使用
	k := 1
	for { //这里等价于for ;;{}
		if k <= 10 {
			fmt.Println("ok", k)
		} else {
			break
		}
		k++
	}

	//字符串遍历方式一：传统方式
	var str string = "hello,world"
	for m := 0; m < len(str); m++ {
		fmt.Printf("%c \n", str[m])
	}

	//字符串遍历方式二：for-range
	var str1 string = "abc"
	for index, val := range str1 {
		fmt.Printf("index = %d,val = %c \n", index, val)
	}

	//字符串遍历方式一：切片
	var str2 string = "hello,world!北京"
	var str3 []rune = []rune(str2)
	for m := 0; m < len(str3); m++ {
		fmt.Printf("%c \n", str3[m])
	}
}
