package main

import (
	"encoding/binary"
	"encoding/json"
	"exec/chat_room/common"
	"fmt"
	"net"
)

//写一个函数，完成登录
func login(userId int, userPwd string) (err error) {

	//定一个协议
	//fmt.Printf("userId = %d,userPwd = %s\n", userId, userPwd)
	//return nil

	//连接服务器
	conn, dial_err := net.Dial("tcp", "localhost:8889")
	if dial_err != nil {
		fmt.Printf("net.Dial err = %v\n", dial_err)
		return dial_err
	}

	defer conn.Close()

	//通过conn发送消息给服务
	var mes common.Message
	mes.Type = common.LoginMesType

	//创建一个LoginMes结构体
	var loginMes common.LoginMes
	loginMes.Userid = userId
	loginMes.UserPwd = userPwd

	//序列化loginMes
	data, json_err := json.Marshal(loginMes)
	if json_err != nil {
		fmt.Printf("json.Marshal err,err = %v", json_err)
		return
	}

	//把data赋给 mes.Data字段
	mes.Data = string(data)

	//将mes序列化
	data, json_err2 := json.Marshal(mes)
	if json_err2 != nil {
		fmt.Printf("json.Marshal err,err = %v", json_err2)
		return
	}

	//此时data为发送的消息
	//先把data长度发送给服务器
	//先获取data长度，转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	//发送长度
	n, conn_byte_err := conn.Write(buf[:4])
	if n != 4 || conn_byte_err != nil {
		return conn_byte_err
	}
	fmt.Printf("客户端，发送长度=%d，内容是=%s", len(data), string(data))

	return
}
