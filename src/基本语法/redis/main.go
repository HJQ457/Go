package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	//连接redis
	conn, err1 := redis.Dial("tcp", "192.168.100.10:6379")
	if err1 != nil {
		fmt.Printf("连接redis失败,err=%v", err1)
	}

	defer conn.Close()

	//登录认证
	_, err2 := conn.Do("auth", "123456")
	if err2 != nil {
		fmt.Printf("登录失败,err=%v", err2)
	}

	//set一个key
	_, err3 := conn.Do("set", "name", "Tom")
	if err3 != nil {
		fmt.Printf("set错误，err=%v", err3)
	}

	//get一个key
	//因为name是string，需要转换
	r, err4 := redis.String(conn.Do("get", "name"))
	if err4 != nil {
		fmt.Printf("get错误，err=%v", err4)
	}

	fmt.Println(r)
}
