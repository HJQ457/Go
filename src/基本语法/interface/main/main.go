package main

import "fmt"

//声明一个接口
type Usb interface {
	Start()
	Stop()
}

type Usb2 interface {
	Start()
	Stop()
}

type Phone struct {
}

type Camera struct {
}

type Computer struct {
}

//让phone实现usb接口方法
func (phone Phone) Start() {
	fmt.Println("手机开始工作")
}

func (phone Phone) Stop() {
	fmt.Println("手机停止工作")
}

//让camera实现usb接口方法
func (camera Camera) Start() {
	fmt.Println("相机开始工作")
}

func (camera Camera) Stop() {
	fmt.Println("相机停止工作")
}

//working方法，接收一个Usb接口类型的变量
//只要实现了Usb接口（所谓实现Usb接口，就是指实现了 Usb接口声明所有方法）
func (computer Computer) Working(usb Usb) {
	//func (computer Computer) Working(usb Usb2) {	也可实现，因为golang中实现接口只基于方法

	//通过Usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}
func main() {
	var computer Computer
	phone := Phone{}
	camera := Camera{}

	//golang中实现接口只基于方法
	computer.Working(phone)
	computer.Working(camera)
}
