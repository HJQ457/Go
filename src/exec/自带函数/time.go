package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("now = %v,now type = %T\n", now, now)

	//通过now获取年月日
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//格式化日期时间
	fmt.Printf("当前年月日 %02d-%02d-%02d %02d:%02d:%02d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	//格式化输出第二种方式(值固定)
	fmt.Printf(now.Format("2006/01/02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("01"))
	fmt.Println()

	//tim()

	//每隔一秒打印一个数字，100截至
	//var i int = 0
	//for  {
	//	i++
	//	fmt.Println(i)
	//	//休眠
	//	//time.Sleep(time.Second)//休眠一秒
	//	time.Sleep(time.Millisecond * 100)//休眠0.1秒
	//	if i == 100 {
	//		break
	//	}
	//}

	//Unix和UnixNano的使用
	fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n", time.Now(), time.Now().UnixNano())

}

func tim() {
	fmt.Printf(time.Now().Format("2006/01/02 15:04:05"))
	fmt.Println()
	fmt.Printf(time.Now().Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(time.Now().Format("15:04:05"))
	fmt.Println()
}
