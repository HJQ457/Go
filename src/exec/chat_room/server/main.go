package main

import (
	"encoding/binary"
	"encoding/json"
	"exec/chat_room/common"
	"fmt"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes common.Message, conn_read_err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据。。。")

	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了conn则就不会阻塞
	_, conn_read_err = conn.Read(buf[:4])
	if conn_read_err != nil {
		//fmt.Printf("conn.Read err=%v", conn_read_err)
		return
	}
	//根据buf[:4]转换成uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	//根据pkgLen读取消息内容
	n, conn_read_buf_err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || conn_read_buf_err != nil {
		//conn_read_buf_err = errors.New("read pkg body error")
		return
	}
	//把pkgLen反序列化 -> message.Message
	Unmarshal_err := json.Unmarshal(buf[:pkgLen], &mes)
	if Unmarshal_err != nil {
		fmt.Println("json.Unmarshal err=", Unmarshal_err)
		return
	}
	return
}

//编写一个函数serverProcessLogin函数，专门处理登录请求
//func serverProcessLogin(conn net.Conn, mes common.Message) (err error) {
//	//先从mes中取出mes.Data，并直接反序列化成LoginMes
//	var loginMes common.Message
//	json.Marshal(mes.Data)
//}

//根据客户端发送消息种类不同，决定调用那个函数处理
func serverProcessMes(conn net.Conn, mes *common.Message) (err error) {
	switch mes.Type {
	case common.LoginMesType:
		//err := serverProcessLogin(conn, mes)
		//if err != nil {
		//	return err
		//}
		//处理登录逻辑
	case common.RegisterMesType:
	//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理。。。")
	}
	return
}

func process(conn net.Conn) {

	defer conn.Close()

	//循环读取客户端发送的信息
	for {
		//读取数据包，封装成一个函数readPkg()，返回Message，Err
		mes, readPkg_err := readPkg(conn)
		if readPkg_err != nil {
			if readPkg_err == io.EOF {
				fmt.Println("客户端退出，服务器端也正常退出。。。")
				return
			} else {
				fmt.Println("readPkg err=", readPkg_err)
				return
			}
		}
		fmt.Println("mes=", mes)
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
