package main

import (
	"fmt"
	reflect "reflect"
)

func reflectTest(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v ，rVal类型为 %T\n", rVal, rVal)

	rType := reflect.TypeOf(b)
	fmt.Printf("rType = %v\n", rType)

	rkind := rVal.Kind()
	fmt.Printf("rKind = %v\n", rkind)

	rIn := reflect.ValueOf(b).Interface()
	fmt.Printf("rInterface = %v\n", rIn)

	rFloat := rIn.(float64)
	fmt.Printf("rFloat = %v\n", rFloat)
}

func reflectChange(b interface{}) {
	reflect.ValueOf(b).Elem().SetFloat(2.1)
}
func main() {
	var v float64 = 1.2
	reflectTest(v)

	reflectChange(&v)
	fmt.Println(v)
}
