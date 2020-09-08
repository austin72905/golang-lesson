package main

import (
	"fmt"
	"reflect"
)

func main() {
	mons := Monster{
		Name:  "dragon",
		Age:   400,
		Score: 40.10,
	}

	//使用反射遍歷struct的屬性，調用struct 的方法，獲取struct 標籤的值
	TestRfctStruct(mons)
}

//Monster struct
//定義一個struct
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

//Print method print some monter info
//Monster 方法1
func (mons Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(mons)
	fmt.Println("---end---")
}

//GetSum method get the sum of two num
//Monster 方法2
func (mons Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//SetVal method set value to monster
//Monster 方法3
func (mons Monster) SetVal(name string, age int, score float32, sex string) {
	mons.Name = name
	mons.Age = age
	mons.Score = score
	mons.Sex = sex
}

//TestRfctStruct reflect struct to get it's value 、property 、tag value
//使用反射遍歷struct的屬性，調用struct 的方法，獲取struct 標籤的值
func TestRfctStruct(val interface{}) {
	rType := reflect.TypeOf(val) //獲取val 的 reflect.Type
	rVal := reflect.ValueOf(val) //獲取val 的 reflect.Value
	rKind := rVal.Kind()         //獲取 val的類別

	//如果傳入不是struct ，跳出
	if rKind != reflect.Struct {
		fmt.Println("it must be a struct")
		return
	}

	//傳入的struct 有幾個屬性
	valPropCount := rVal.NumField()
	fmt.Printf("struct has %d fields\n", valPropCount)

	//遍歷strcut 所有屬性
	for i := 0; i < valPropCount; i++ {
		fmt.Printf("Field %d: 值為 = %v\n", i, rVal.Field(i)) //返回傳入struct 的第 i 個屬性 ，返回一個reflect.Value ，要運算要再型轉
		//注意這邊是用type
		tagVal := rType.Field(i).Tag.Get("json") //Type裡有個方法Field(i int) StructField   =>StructField是一個結構體
		//Get() 裡面的值可以自定義

		//不是所有屬性都有tag
		if tagVal != "" {
			fmt.Printf("Field %d的tag為 = %v\n", i, tagVal)
		}
	}

	//調用傳入struct 的方法
	//傳入的struct 有幾個方法
	methodCount := rVal.NumMethod()
	fmt.Printf("struct has %d methods\n", methodCount)

	//調用方法
	rVal.Method(1).Call(nil) //調用傳入struct 的第2個方法 (字典序)
	//所以其實調用的是Print()

	//如果要調用有參數的方法，參數類型是[]slice

	//定義參數
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10)) //將常數10 轉成reflect.Value 並放入切片
	params = append(params, reflect.ValueOf(40))

	//調用Monster第1個方法
	result := rVal.Method(0).Call(params) //Call 返回是一個 reflect.Value類型
	fmt.Println("result=", result[0].Int())

}

//使用反射遍歷struct的屬性，調用struct 的方法，獲取struct 標籤的值

//func (v Value) Field(i int) Value
//返回傳入struct 的第 i 個屬性

//func (v Value) Method(i int) Value
//可以獲得傳入的struct 本身的第i個方法

//func (v Value) Call(in []Value) []Value   =>傳入的參數是[]reflect.Value
//呼叫 用Method 獲得的方法
