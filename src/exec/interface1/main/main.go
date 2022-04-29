package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

//声明一个Hero结构体切片类型
type HeroSlice []Hero

//实现interface接口

func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}

func main() {

	//先定义一个数组/切片
	var intSlice = []int{0, -1, 10, 7, 90}

	//冒泡排序
	for i := 0; i < len(intSlice)-1; i++ {
		var temp int
		if intSlice[i] > intSlice[i+1] {
			temp = intSlice[i]
			intSlice[i] = intSlice[i+1]
			intSlice[i+1] = temp
		}
	}
	fmt.Println(intSlice)

	//系统提供的方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}

	//排序前顺序
	for _, v := range heros {
		fmt.Println(v)
	}

	fmt.Println()

	//排序后顺序
	sort.Sort(heros)
	for _, v := range heros {
		fmt.Println(v)
	}
}
