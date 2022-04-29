package main

import "fmt"

type Monkey struct {
	Name string
}

//声明接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树")
}

type LittleMonkey struct {
	Monkey //继承
}

func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习，会飞翔")
}

func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "通过学习，会游泳")
}

func main() {
	var monkey LittleMonkey
	monkey.Name = "wukong"
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
