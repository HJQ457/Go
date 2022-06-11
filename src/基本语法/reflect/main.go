package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {

	//获取类型
	rType := reflect.TypeOf(b)
	fmt.Printf("rType=%v\n", rType)

	//获取值
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v\n", rVal)

	fmt.Printf("rVal=%v,rVal type=%T\n", rVal, rVal)

	//将rVal转成interface
	iV := rVal.Interface()
	//将interface{}通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Printf("num2=%v,num2 type=%T\n", num2, num2)
}

//对结构体反射
func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Printf("rType=%v\n", rType)

	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v\n", rVal)

	fmt.Printf("rVal=%v,rVal type=%T\n", rVal, rVal)

	iV := reflect.ValueOf(b).Interface()
	num3 := iV.(student)
	fmt.Printf("num3=%v,num3 type=%T\n", num3, num3)

	//获取变量对应的kind
	kind1 := rVal.Kind()
	kind2 := rType.Kind()
	fmt.Printf("kind1=%v,kind2=%v\n", kind1, kind2)

	//获取字段个数
	num := reflect.ValueOf(b).NumField()
	fmt.Printf("结构体有 %v 个字段\n", num)

	//遍历字段
	//for i := 0; i < num; i++ {
	//	fmt.Printf("第%v的字段值为%v\n", i+1, reflect.ValueOf(b).Field(i))
	//	//获取标签
	//	fmt.Printf("第%v的字段标签为%\n", i+1, reflect.TypeOf(b).Field(i).Tag.Get("json"))
	//}

	//获取结构体方法个数
	numMethod := reflect.ValueOf(b).NumMethod()
	fmt.Printf("结构体有%v个方法", numMethod)

	reflect.ValueOf(b).Method(0).Call(nil)

}

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (s student) Print() {
	fmt.Println("------start------")
	fmt.Printf(s.Name)
	fmt.Printf("------end------")
}

func (s student) Set(n1, n2 int) int {
	return n1 + n2
}

func main() {
	var num int = 100
	reflectTest01(num)

	stu := student{
		Name: "Tom",
		Age:  10,
	}

	fmt.Println()

	reflectTest02(stu)
}
