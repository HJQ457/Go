package main

import "fmt"

func main() {
	var i int = 1
	for {
		if i > 10 {
			break
		}
		fmt.Println("hello world", i)
		i++
	}
}
