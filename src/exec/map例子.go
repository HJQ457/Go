package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string) {

	//判断user中是否有name
	if users[name] != nil {
		users[name]["pwd"] = "88888888"
	} else {
		users[name] = make(map[string]string)
		users[name]["pwd"] = "88888888"
		users[name]["nickname"] = "昵称" + name
	}
}
func main() {
	users := make(map[string]map[string]string, 10)
	modifyUser(users, "tom")
	modifyUser(users, "marry")

	fmt.Println(users)
}
