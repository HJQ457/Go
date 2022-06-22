package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	dial, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("client dial err=%v", err)
		return
	}

	for {
		//客户端发送单行单据，然后退出
		reader := bufio.NewReader(os.Stdin) //os.stdin代表标准输入

		//从终端读取一行用户输入，并准备发送给服务器
		line, err2 := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err2)
		}

		//如果用户输入exit就退出
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出...")
			break
		}

		//再将line发送给服务器
		_, err2 = dial.Write([]byte(line))
		if err2 != nil {
			fmt.Printf("conn.write err=", err2)
		}
		//fmt.Printf("客户端发送了%d字节的数据，并退出\n", n)
	}
}
