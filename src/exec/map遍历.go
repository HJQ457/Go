package main

import "fmt"

func main() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"

	for k, v := range cities {
		fmt.Printf("k = %v,v = %v\n", k, v)
	}

	studentMap := make(map[string]map[string]string)

	studentMap["stu1"] = make(map[string]string, 2)
	studentMap["stu1"]["name"] = "tom"
	studentMap["stu1"]["sex"] = "男"

	studentMap["stu2"] = make(map[string]string, 3)
	studentMap["stu2"]["name"] = "tom"
	studentMap["stu2"]["sex"] = "女"
	studentMap["stu2"]["address"] = "北京长安街"

	for k1, v1 := range studentMap {
		fmt.Println("k1 = ", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2 = %v,v2 = %v\n", k2, v2)
		}
		fmt.Println()
	}
}
