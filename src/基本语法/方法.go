package main

import "fmt"

type Per struct {
	name string
}

//给A结构体绑定一个方法
func (a Per) test() {
	a.name = "jack"
	fmt.Println("test():", a.name)
}

func (a Per) speak() {
	fmt.Printf("speak(): %v是一个好人", a.name)
}

func (a Per) jisuan(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res = res + i
	}
	fmt.Println("jisuan():", res)
}

func main() {
	var p Per
	p.name = "tom"
	p.test() //调用方法
	fmt.Println("main():", p.name)
	p.speak()
	p.jisuan(100)
}
