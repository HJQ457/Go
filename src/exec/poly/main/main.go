package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

type Computer struct {
}

func (phone Phone) Start() {
	fmt.Println("手机开始工作")
}

func (phone Phone) Stop() {
	fmt.Println("手机停止工作")
}

func (p Phone) Call() {
	fmt.Println("手机开始打电话")
}

func (camera Camera) Start() {
	fmt.Println("相机开始工作")
}

func (camera Camera) Stop() {
	fmt.Println("相机停止工作")
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	//usb.Call() //error
	//需要类型断言
	if phone, ok := usb.(Phone); ok == true {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	//var computer Computer
	//var camera Camera
	//computer.Working(Phone{})
	//computer.Working(camera)

	//var usbArr []Usb = make([]Usb, 5)
	var usbArr [3]Usb
	usbArr[0] = Phone{"华为"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"索尼"}
	fmt.Println(usbArr)

	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
	}
}
