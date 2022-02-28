package main

import "fmt"

func main() {
	cites := make(map[string]string)
	cites["no1"] = "北京"
	cites["no2"] = "天津"
	cites["no3"] = "上海"

	fmt.Println(cites)

	//因为no3已经存在，所以下面是修改
	cites["no3"] = "深圳"
	fmt.Println(cites)

	//删除
	delete(cites, "no1")
	fmt.Println(cites)

	//删除不存在的key时，删除不会操作，也不会报错
	delete(cites, "no4")
	fmt.Println(cites)

	//一次性删除所有的key
	//1.遍历所有的key删除
	//2.直接make新的空间，被gc回收
	cites = make(map[string]string)
	fmt.Println(cites)

	//查找
	val, ok := cites["no2"]
	if ok {
		fmt.Printf("有no2 key值为%v\n", val)
	} else {
		fmt.Printf("没有no2 key\n")
	}
}
