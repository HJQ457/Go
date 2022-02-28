package main

import "fmt"

func main() {

	//第一种方式
	var a map[string]string
	//在使用map前，需要先make，makde的作用是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no1"] = "武松"
	a["no3"] = "吴用"

	fmt.Println(a)

	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"

	fmt.Println(cities)

	//第三种方式
	heroes := map[string]string{
		"hero1": "宋江",
		"hero2": "卢俊义",
	}
	heroes["hero3"] = "林冲"

	fmt.Println(heroes)
}
