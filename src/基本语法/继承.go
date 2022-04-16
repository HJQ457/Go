package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float64
}

//两个结构体相同的方法
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名字：%v，年龄：%v，分数：%v", stu.Name, stu.Age, stu.Score)
}

//两个结构体相同的方法
func (stu *Student) SetScore(score float64) {
	stu.Score = score
}

type Pupil struct {
	Student //嵌入匿名结构体
}

//特有方法
func (p *Pupil) Testing() {
	fmt.Println("小学生正在考试")
}

type Graduate struct {
	Student
}

//特有方法
func (p *Graduate) Testing() {
	fmt.Println("大学生正在考试")
}

func main() {
	var p Pupil
	p.Name = "Tom"
	p.SetScore(60)
	p.Testing()
	p.ShowInfo()

	fmt.Println()

	graduate := &Graduate{}
	graduate.Student.Name = "jack"
	graduate.ShowInfo()
}
