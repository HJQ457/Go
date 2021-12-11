package main

func main() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue
			}
			println(j)
		}
	}

	for m := 0; m < 4; m++ {
		for n := 0; n < 10; n++ {
			if n == 2 {
				continue
			}
			println(n)
		}
	}
}
