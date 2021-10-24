package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "true"
	var b bool

	//b ,_ = strconv.ParseBool(str) 本身会返回两个值（value bool,err error）,因只想获取value bool,所以用_忽略err
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T,b=%v\n", b, b)

	var str2 string = "1234590"
	var n1 int64
	var n2 int

	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)

	fmt.Printf("n2 type %T,n2=%d\n", n2, n2)

	var str3 string = "123.456"
	var f1 float64

	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 tpye %T,f1=%f\n", f1, f1)

	//string 转换基本类型要确定能否转成有效的数据，如果转不成有效数据，则直接转换成默认值
	var str4 string = "hello"
	var n3 int64
	n3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type %T,n3=%d\n", n3, n3)
}
