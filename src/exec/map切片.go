package main

import "fmt"

func main() {
	var monsters []map[string]string = make([]map[string]string, 2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "Tom"
		monsters[0]["age"] = "500"
	}
	fmt.Println(monsters)

	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "Sam"
		monsters[1]["age"] = "50"
	}
	fmt.Println(monsters)

	//先定义monster信息
	newMonster := map[string]string{
		"name": "Jack",
		"age":  "30",
	}
	//用append函数，动态增加monster
	monsters = append(monsters, newMonster)
}
