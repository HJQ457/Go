package main

import "fmt"

func main() {
	studentMap := make(map[string]map[string]string)

	studentMap["stu1"] = make(map[string]string, 2)
	studentMap["stu1"]["name"] = "tom"
	studentMap["stu1"]["sex"] = "男"

	studentMap["stu2"] = make(map[string]string, 3)
	studentMap["stu2"]["name"] = "tom"
	studentMap["stu2"]["sex"] = "女"
	studentMap["stu2"]["address"] = "北京长安街"

	fmt.Println(studentMap)
}
