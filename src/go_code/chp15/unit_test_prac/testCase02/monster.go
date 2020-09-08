package testcase02

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//Monster struct
type Monster struct {
	Name  string
	Age   int
	Skill string
}

//Store method belong to Monster Serialize struct to json string
//monster 方法 store ，將monster變數，序列化保存到文件
func (mons *Monster) Store() bool {
	//1. 序列化
	data, err := json.Marshal(mons)
	if err != nil {
		fmt.Println("序列化失敗")
		return false
	}
	//2. 保存到文件
	filePath := "C:/Users/USER/Desktop/文件檔/unit_test"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err =", err)
		return false
	}

	return true
}

//Retore method belong to Monster Deserialize json string to struct
//monster 反序列化
func (mons *Monster) Retore() bool {
	//1. 從文件讀取序列化的字串
	filePath := "C:/Users/USER/Desktop/文件檔/unit_test"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file err =", err)
		return false
	}

	//2. 反序列化成物件
	err = json.Unmarshal(data, mons)
	if err != nil {
		fmt.Println("undeSerialize err =", err)
		return false
	}

	return true
}
