package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	testStruct()
	testMap()
	testSlice()
}

func testStruct() {
	//將struct序列化
	stud := Student{
		Name: "austin",
		Age:  15,
	}
	//序列化
	stuJsStr, err := json.Marshal(&stud)
	if err != nil {
		fmt.Printf("err occured when serialize errMsg:%v\n", err)
	}

	fmt.Printf("student = %v\n", string(stuJsStr)) //是回傳[]byte，要在轉字串

}

func testMap() {

	var testMap map[string]interface{}     //value 是空介面，可接任何值
	testMap = make(map[string]interface{}) //要先make不然抱錯
	testMap["name"] = "jack"
	testMap["age"] = 16

	//序列化
	mapJsStr, err := json.Marshal(testMap) //本身就是 指針類型 不用&
	if err != nil {
		fmt.Printf("err occured when serialize errMsg:%v\n", err)
	}

	fmt.Printf("student1 = %v\n", string(mapJsStr)) //是回傳[]byte，要在轉字串

}

func testSlice() {
	var testSlice []map[string]interface{}
	var map1 map[string]interface{}
	map1 = make(map[string]interface{}) //要先make不然抱錯
	map1["name"] = "tom"
	map1["age"] = 17

	var map2 map[string]interface{}
	map2 = make(map[string]interface{}) //要先make不然抱錯
	map2["name"] = "mary"
	map2["age"] = 17

	//將map 加入切片
	testSlice = append(testSlice, map1)
	testSlice = append(testSlice, map2)

	//序列化
	silceJsStr, err := json.Marshal(testSlice) //本身就是 指針類型 不用&
	if err != nil {
		fmt.Printf("err occured when serialize errMsg:%v\n", err)
	}

	fmt.Printf("student1 = %v\n", string(silceJsStr)) //是回傳[]byte，要在轉字串
}

//裡面的屬性要"大寫"才有辦法使用json.Marshal
//可以在屬性後面加tag ，指定序列化後的key名
//內部是使用反射機制實現

//Student struct
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//json

//在js 裡面 所有的一切都是對象

//序列化: 將有key value 形式的數據(map、slice、struct、arr都可以)轉為 json 字串

//json 包 func Marshal(v interface{}) ([]byte, error)

//對基本數據 序列化，只是把它變字串，沒有key
