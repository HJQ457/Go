package main

import "fmt"

func BinaryFind(slice *[]int, leftIndex int, rightIndex int, findval int) {

	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	middle := (leftIndex + rightIndex) / 2

	if (*slice)[middle] > findval {
		BinaryFind(slice, leftIndex, middle-1, findval)
	} else if (*slice)[middle] < findval {
		BinaryFind(slice, middle+1, rightIndex, findval)
	} else {
		fmt.Printf("找到了，下标为%v \n", middle)
	}

}

func main() {
	slice := []int{1, 8, 10, 89, 1000, 1234}
	BinaryFind(&slice, 0, len(slice)-1, 400)
}
