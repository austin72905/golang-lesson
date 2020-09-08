package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	testDeserialize()
	testDsMap()
	testDsSlice()
}

func testDeserialize() {
	jsStr := "{\"Name\":\"austin\",\"Age\":15}"

	//new 一個 struct
	var stud Student

	err := json.Unmarshal([]byte(jsStr), &stud)
	if err != nil {
		fmt.Printf("err occured when deserialize ... errMsg=%v\n", err)

	}

	fmt.Printf("stud=%v\n", stud)
}

func testDsMap() {
	jsStr := "{\"age\":16,\"name\":\"jack\"}"

	//反序列化的時候不用make， json.Unmarshal 底層有幫我們做了
	var testMap map[string]interface{}

	//反序列化
	err := json.Unmarshal([]byte(jsStr), &testMap)
	if err != nil {
		fmt.Printf("err occured when deserialize ... errMsg=%v\n", err)

	}

	fmt.Printf("testMap=%v\n", testMap)
}

func testDsSlice() {
	var testSicle []map[string]interface{}

	//反序列化的時候不用make， json.Unmarshal 底層有幫我們做了
	jsStr := "[{\"age\":17,\"name\":\"tom\"},{\"age\":17,\"name\":\"mary\"}]"
	//反序列化
	err := json.Unmarshal([]byte(jsStr), &testSicle)
	if err != nil {
		fmt.Printf("err occured when deserialize ... errMsg=%v\n", err)

	}

	fmt.Printf("testSicle=%v\n", testSicle)
}

//Student struct
type Student struct {
	Name string
	Age  int
}

//反序列化

//反序列化的時候不用make， json.Unmarshal 底層有幫我們做了

//如果json字串 是由程序獲得的，不需要再用 \ 對"作轉譯處理
