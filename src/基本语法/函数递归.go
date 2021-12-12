package main

import "fmt"

func main() {
	tests(4)
}

func tests(n int) {
	if n > 2 {
		n--
		tests(n)
	}
	fmt.Println(n)
}
