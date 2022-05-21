package monster

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

//序列化
func (monster *Monster) Store() bool {
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败")
		return false
	}

	//保存文件
	filePath := "D:\\gocode\\src\\exec\\test_case\\monster.txt"

	data1, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("open file err , %v", err)
	}

	defer data1.Close()

	writer := bufio.NewWriter(data1)
	_, err1 := writer.WriteString(string(data))
	if err1 != nil {
		fmt.Printf("writer file err %v", err1)
	}

	writer.Flush()

	return true
}

//反序列化
func (monster *Monster) Restore() bool {
	filePath := "D:\\gocode\\src\\exec\\test_case\\monster.txt"

	data2, err2 := ioutil.ReadFile(filePath)
	if err2 != nil {
		fmt.Printf("read file err %v", err2)
		return false
	}

	err3 := json.Unmarshal(data2, monster)
	if err3 != nil {
		fmt.Printf("unmarshal err %v", err3)
		return false
	}

	return true
}
