package main

import (
	"exec/init/utils"
	"fmt"
)

//执行顺序：变量>init()>main()
func main() {
	fmt.Println("Age = ", utils.Age, "Name = ", utils.Name)
}
