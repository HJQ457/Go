package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {

	//循环接收数据
	defer conn.Close()

	for {
		fmt.Printf("服务器在等待客户端%s发送信息", conn.RemoteAddr().String())
		buf := make([]byte, 1024)
		read, err3 := conn.Read(buf)
		if err3 == io.EOF {
			fmt.Println("客户端退出")
			return
		}

		//显示客户端发送的内容到服务器终端
		fmt.Print(string(buf[:read]))
	}
}
func main() {
	fmt.Println("服务开始监听。。。")
	//net.Listen("tcp", "127.0.0.1:8888") //只支持ipv4
	listen, err := net.Listen("tcp", "0.0.0.0:8888") //支持ipv4和ipv6
	if err != nil {
		fmt.Printf("listen err=%v", err)
		return
	}

	defer listen.Close()

	//等待客户端连接
	for {
		fmt.Println("等待客户端连接。。。")
		conn, err2 := listen.Accept()
		if err2 != nil {
			fmt.Printf("Accept() err=%v", err2)
		} else {
			fmt.Printf("Accept() success conn=%v，客户端ip为%v", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}

	fmt.Printf("listen=%v", listen)
}
