package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义一个全局pool
var pool *redis.Pool

//启动程序时，初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //表示和数据库最大连接数，0为无限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化连接的代码，连接哪个ip
			password := redis.DialPassword("123456")
			database := redis.DialDatabase(1)
			return redis.Dial("tcp", "192.168.100.10:6379", password, database)
		},
	}
}

func main() {
	//先从pool取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_, err1 := conn.Do("set", "name", "Tom")
	if err1 != nil {
		fmt.Printf("conn.Do err = ", err1)
		return
	}

	//取出
	r2, err2 := redis.String(conn.Do("get", "name"))
	if err2 != nil {
		fmt.Printf("get err = ", err2)
		return
	}
	fmt.Println(r2)

	pool.Close() //关闭连接池，以下代码会报错

	conn2 := pool.Get()

	_, err4 := conn2.Do("set", "name1", "Jack")
	if err4 != nil {
		fmt.Printf("conn.Do err = ", err4)
		return
	}

	r3, err3 := redis.String(conn2.Do("get", "name1"))
	if err3 != nil {
		fmt.Printf("get err = ", err3)
		return
	}
	fmt.Println(r3)
}
