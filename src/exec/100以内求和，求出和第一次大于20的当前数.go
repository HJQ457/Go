package main

func main() {
	sum := 0
	i := 1

	for ; i <= 100; i++ {
		sum = sum + i
		if sum > 20 {
			break
		}
	}

	println(i)
}
