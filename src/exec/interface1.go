package main

import "fmt"

type Binterface interface {
	test01()
}

type T interface {
}

type Cinterface interface {
	test02()
}

type Ainterface interface {
	Binterface
	Cinterface
	test03()
}

type Stu struct {
}

func (stu Stu) test01() {

}

func (stu Stu) test02() {

}

func (stu Stu) test03() {

}

func main() {
	var stu Stu
	var a Ainterface = stu
	a.test01()

	var t T = stu //空接口也可以使用
	fmt.Println(t)
}
