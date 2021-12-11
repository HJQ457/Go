package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int = 0
	rand.Seed(time.Now().UnixNano()) //返回从1970:01:01的00:00:00到现在的纳秒数

	for {
		//rand.Seed(time.Now().Unix())	//返回从1970:01:01的00:00:00到现在的秒数

		n := rand.Intn(100) + 1 //[0,100)
		count++
		if n == 99 {
			break
		}
	}

	fmt.Printf("生成99一共用了%v次", count)

}
