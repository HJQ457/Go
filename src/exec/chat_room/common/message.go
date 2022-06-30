package common

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的类型
}

//定义两个消息
type LoginMes struct {
	Userid   int    `json:"userid"`    //用户id
	UserPwd  string `json:"user_pwd"`  //用户密码
	UserName string `json:"user_name"` //用户名
}

type LoginResMes struct {
	Code  int    `json:"code"`  //返回状态码	500表示用户未注册  200表示登录成功
	Error string `json:"error"` //返回错误信息
}
