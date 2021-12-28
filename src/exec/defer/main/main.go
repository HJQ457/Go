package main

import (
	"exec/defer/utils"
	"fmt"
)

func main() {
	res := utils.Sum(10, 20)
	fmt.Println("res = ", res)
}
