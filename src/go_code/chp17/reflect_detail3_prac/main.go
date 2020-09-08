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
	fmt.Println("修改前 =", mons)
	//使用反射修改struct屬性的值(要傳地址進去)
	TestRfctStruct(&mons)
	fmt.Println("修改後 =", mons)
}

//Monster struct
//定義一個struct
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

//TestRfctStruct func set new value to monster
func TestRfctStruct(val interface{}) {
	//rType := reflect.TypeOf(val)
	rVal := reflect.ValueOf(val)
	rKind := rVal.Kind()
	if rKind != reflect.Ptr && rVal.Elem().Kind() == reflect.Struct {
		fmt.Println("arguement must be struct")
		return
	}

	propCount := rVal.Elem().NumField()
	//修改第一個屬性的值
	rVal.Elem().Field(0).SetString("kibana")
	//打印傳入的struct 有幾個屬性
	for i := 0; i < propCount; i++ {
		fmt.Printf("%d %v\n", i, rVal.Elem().Field(i).Kind())
	}
	fmt.Printf("struct has %d field \n", propCount)
}

//使用反射，修改struct的值
