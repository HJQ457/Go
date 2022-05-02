package main

import "fmt"

type Student struct {
}

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型,值是%v\n", index+1, x)
		case float32:
			fmt.Printf("第%v个参数是float32类型,值是%v\n", index+1, x)
		case float64:
			fmt.Printf("第%v个参数是float64类型,值是%v\n", index+1, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是int类型,值是%v\n", index+1, x)
		case string:
			fmt.Printf("第%v个参数是string类型,值是%v\n", index+1, x)
		case Student:
			fmt.Printf("第%v个参数是Student类型,值是%v\n", index+1, x)
		case *Student:
			fmt.Printf("第%v个参数是*Student类型,值是%v\n", index+1, x)
		default:
			fmt.Println("第%v个参数是 类型 不确定,值是%v\n", index+1, x)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "Tom"
	address := "北京"
	n4 := 300
	n5 := Student{}
	n6 := &Student{}

	TypeJudge(n1, n2, n3, name, address, n4, n5, n6)
}
