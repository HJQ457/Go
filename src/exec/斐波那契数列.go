package main

import "fmt"

func main() {
	fmt.Println(fbn(6))
}

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}
