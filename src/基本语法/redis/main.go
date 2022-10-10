package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	//登录认证方式一
	password := redis.DialPassword("123456")
	database := redis.DialDatabase(0)

	//连接redis
	conn, err1 := redis.Dial("tcp", "192.168.100.10:6379", password, database)
	if err1 != nil {
		fmt.Printf("连接redis失败,err=%v\n", err1)
	}

	defer conn.Close()

	//登录认证方式二
	//_, err2 := conn.Do("auth", "123456")
	//if err2 != nil {
	//	fmt.Printf("登录失败,err=%v\n", err2)
	//}

	//set一个key
	_, err3 := conn.Do("set", "name", "Tom")
	if err3 != nil {
		fmt.Printf("set错误，err=%v\n", err3)
	}

	//get一个key
	//因为name是string，需要转换
	r, err4 := redis.String(conn.Do("get", "name"))
	if err4 != nil {
		fmt.Printf("get错误，err=%v\n", err4)
	}

	fmt.Println("key=", r)

	//set一个Hashkey
	_, err5 := conn.Do("hset", "user01", "name", "John")
	if err5 != nil {
		fmt.Printf("hset错误，err=%v\n", err5)
	}

	_, err10 := conn.Do("hset", "user01", "age", "10")
	if err10 != nil {
		fmt.Printf("hset错误，err=%v\n", err10)
	}

	//get一个Hashkey
	//因为name是string，需要转换
	r, err6 := redis.String(conn.Do("hget", "user01", "name"))
	if err6 != nil {
		fmt.Printf("hget错误，err=%v\n", err6)
	}

	fmt.Println("hget=", r)

	hg, err12 := redis.Strings(conn.Do("hgetall", "user01"))
	if err12 != nil {
		fmt.Printf("hgetall错误，err=%v\n", err12)
	}

	fmt.Println("hgetall=", hg)

	//set一个hmkey
	_, err7 := conn.Do("hmset", "user02", "name", "Jack", "age", "20")
	if err7 != nil {
		fmt.Printf("hmset错误，err=%v\n", err7)
	}

	//get一个hmkey
	hm, err11 := redis.Strings(conn.Do("hmget", "user02", "name", "age"))
	if err11 != nil {
		fmt.Printf("hmget错误，err=%v\n", err11)
	}

	//fmt.Println("hmget=", hm)
	for k, v := range hm {
		fmt.Printf("hm[%v]=%v\n", k, v)
	}

	for i := 0; i < len(hm); i++ {
		fmt.Printf("hm[%v]=%v\n", i, hm[i])
	}

	//获取所有key
	values, value_err := redis.Values(conn.Do("keys", "*"))
	if value_err != nil {
		return
	}

	for i := 0; i < len(values); i++ {
		fmt.Println(string(values[i].([]byte)))
	}

	for _, v := range values {
		fmt.Printf(string(v.([]byte)))
	}
}
