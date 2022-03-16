package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	var keys []int
	for k1, _ := range map1 {
		keys = append(keys, k1)
	}

	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	//遍历切片
	for _, i := range keys {
		fmt.Printf("key = %v,value = %v\n", i, map1[i])
	}
}
