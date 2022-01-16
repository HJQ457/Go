package main

import "fmt"

func main() {
	//层数
	var totalLevel int = 5

	for i := 1; i <= totalLevel; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()

	for n := 1; n <= 5; n++ {

		//打印*号前打印空格
		for k := 1; k <= 5-n; k++ {
			fmt.Print(" ")
		}
		for m := 1; m <= 2*n-1; m++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
