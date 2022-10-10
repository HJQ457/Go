package main

import "fmt"

func main() {
	//parseInt, err := strconv.ParseInt("1000", 10, 0)
	//if err != nil {
	//	return
	//}
	//fmt.Println(parseInt)
	//
	//atoi, err := strconv.Atoi("10")
	//if err != nil {
	//	return
	//}
	//fmt.Println(atoi)
	//
	//formatInt := strconv.FormatInt(100, 10)
	//fmt.Println(formatInt)

	var a []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//b := []int{0, 1, 2}
	b := make([]int, 10)
	copy(b, a)
	fmt.Println(a)
	fmt.Println(b)

	str := "你好，世界！hello world！"
	runes := []rune(str)
	fmt.Println(string(runes[3]))
	fmt.Println(string(runes[11]))
}
