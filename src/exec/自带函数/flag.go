package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var n = flag.Bool("n", false, "This is bool")
	var sep = flag.String("s", " ", "This is string")

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Println(feibo(5))
}

func feibo(n int) []int {

	feibo := make([]int, 10)
	feibo[0] = 1
	feibo[1] = 1

	for i := 2; i <= n; i++ {
		feibo[i] = feibo[i-1] + feibo[i-2]
	}
	return feibo
}
