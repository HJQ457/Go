package main

import "fmt"

type Ainterface interface {
	Say()
}

type Stu struct {
	name string
}

type interger int

func (stu Stu) Say() {
	fmt.Println("stu say方法")
}

func (i interger) Say() {
	fmt.Println("interger say i=", i)
}

func main() {
	var stu Stu
	var a Ainterface = stu
	a.Say()

	var i interger = 10
	var b Ainterface = i
	b.Say()
}
