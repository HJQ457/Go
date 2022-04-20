package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//1.统计字符串长度
	str := "hello北"
	//golang字符串编码统一为utf-8,len函数获取的是字节，不是字符
	fmt.Println(len(str))

	str2 := "hello北京"

	//2.字符串遍历，处理中文问题 r := []rune(str)
	str3 := []rune(str2)
	for i := 0; i < len(str3); i++ {
		fmt.Printf("字符 = %c\n", str3[i])
	}

	//3.字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换的结果：", n)
	}

	//4.整数转字符串
	str4 := strconv.Itoa(12345)
	fmt.Printf("str4 = %v,str = %T\n", str4, str4)

	//5.字符串转byte
	bytes := []byte("hello world")
	fmt.Println(bytes)

	//6.byte转字符串
	str5 := string([]byte{97, 98, 99})
	fmt.Printf("str5 = %v\n", str5)

	//7.10进制转换2，8，16
	str6 := strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是：%v\n", str6)
	str7 := strconv.FormatInt(123, 8)
	fmt.Printf("123对应的八进制是：%v\n", str7)
	str8 := strconv.FormatInt(123, 16)
	fmt.Printf("123对应的十六进制是：%v\n", str8)

	//8.查找子串是否在指定的字符串中
	str9 := strings.Contains("seafood", "fod")
	fmt.Printf("fod是否在seafood中：%v\n", str9)

	//9.统计一个字符串中有几个指定的字串
	num1 := strings.Count("ceheese", "e")
	fmt.Printf("e出现的次数为：%v\n", num1)

	//10.不区分大小写比较字符串（==是区分大小写的）
	b := strings.EqualFold("abc", "Abc")
	fmt.Printf("abc和Abc是否一样：%v\n", b)

	//11.返回字串在字符串中第一次出现的index值，如果没有会返回-1
	index := strings.Index("NLT_abc", "abc")
	fmt.Printf("索引值为：%v\n", index)

	//12.返回字串在字符串中最后一次出现的index值，如果没有会返回-1
	index1 := strings.LastIndex("go_golang", "go")
	fmt.Printf("go出现的最后一次为：%v\n", index1)

	//13.将指定的字符串替换成另外一个子串，n指定替换几个，如果n=-1全部替换
	str10 := strings.Replace("go go go hello", "go", "go语言", 2)
	fmt.Printf("替换后的字符串：%v\n", str10)

	//14.按照指定的字符，为分割标识，将一个字符串拆分成字符串数组
	strArr := strings.Split("hello,world,ok", ",")
	fmt.Printf("strArr = %v\n", strArr)

	//15.将字符串进行大小写转换
	str11 := "golang Hello"
	str12 := strings.ToLower(str11)
	fmt.Printf("转换小写后：%v\n", str12)
	str13 := strings.ToUpper(str11)
	fmt.Printf("转换大写后：%v\n", str13)

	//16.去掉字符串左右两边空格
	str14 := strings.TrimSpace(" tn a lone go  ")
	fmt.Printf("去掉空格后=%q\n", str14)

	//17.将字符串左右两边指定的字符去
	str15 := strings.Trim("! hello! ", " !")
	fmt.Printf("删除指定字符：%q\n", str15)

	//18.判断字符串是否以指定的字符串开头
	b1 := strings.HasPrefix("ftp://92.168.0.1", "ftp")
	fmt.Printf("是否以ftp开头：%v\n", b1)

}
