package main

func addUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func getSub(n int, m int) int {
	res := n + m
	return res
}
