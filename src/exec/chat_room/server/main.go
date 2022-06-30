package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {

	defer conn.Close()

	//循环读取客户端发送的信息
	for {
		buf := make([]byte, 8096)
		fmt.Println("读取客户端发送的数据。。。")
		n, conn_read_err := conn.Read(buf[:4])
		if n != 4 || conn_read_err != nil {
			fmt.Printf("conn.Read err=%v", conn_read_err)
			return
		}
		fmt.Printf("读取到的buf=%v\n", buf[:4])
	}
}

func main() {
	//提示信息
	fmt.Println("服务器在6379端口监听。。。")
	listen, err1 := net.Listen("tcp", "0.0.0.0:8889")
	if err1 != nil {
		fmt.Printf("服务器listen失败,err=%v\n", err1)
		return
	}
	defer listen.Close()

	//一旦监听成功，等待客户端连接服务器
	for {
		fmt.Println("等待客户端连接。。。")
		conn, err2 := listen.Accept()
		if err2 != nil {
			fmt.Printf("Accept() err=%v\n", err2)
			return
		}

		//一旦连接成功，则启动一个协程
		go process(conn)
	}
}
