package main

import "fmt"

func main() {
	var n int = 6

	for i := 0; i <= n; i++ {
		fmt.Printf("%v + %v = %v \n", i, n-i, n)
	}
}
