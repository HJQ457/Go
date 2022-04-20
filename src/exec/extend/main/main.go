package main

import "exec/extend/model"

func main() {
	var c model.C
	//c.Name = "Tom" //error
	c.A.Name = "Tom"

	var d model.D
	//d.Name //报错。有名结构体不会向上找
	d.A.Name = "Jack"
}
