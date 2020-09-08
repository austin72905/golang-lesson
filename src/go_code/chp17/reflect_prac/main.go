package main

import (
	"fmt"
	"reflect"
)

func main() {
	//反射轉換基本類型
	// var num int = 100
	// reflectTest01(num)

	//反射轉換結構體
	stud := Student{
		Name: "austin",
		Age:  "10",
	}
	reflectTest02(stud)
}

//反射轉換基本類型
func reflectTest01(val interface{}) {

	//通過反射或取傳入變數的type、kind、值
	//1. 獲取 reflect.Type
	rType := reflect.TypeOf(val)
	fmt.Println("rType=", rType) //int  但他的類型其實是reflect.Type，只是它可以調用其他函數

	//2.獲取 reflect.Value
	rVal := reflect.ValueOf(val)

	fmt.Printf("rVal=%v rVal type= %T\n", rVal, rVal) //100  reflect.Value
	//雖然打印100 但他實際類型還是reflect.Value
	//如果要運算，可以使用reflect.Value 裡類型轉換的方法
	//如果傳進來是int 就不能用Float()，數據類型要匹配
	numPlus := 32 + rVal.Int()
	fmt.Println("numPlus =", numPlus)

	//將reflect.Value 轉成interface{}
	iVal := rVal.Interface()
	//通過類型斷言轉成想要的類型
	num := iVal.(int)
	fmt.Println("num=", num)
}

//反射轉換結構體
func reflectTest02(val interface{}) {

	//通過反射或取傳入變數的type、kind、值
	//1. 獲取 reflect.Type
	rType := reflect.TypeOf(val)
	fmt.Println("rType=", rType) //  但他的類型其實是reflect.Type，只是它可以調用其他函數

	//2.獲取 reflect.Value
	rVal := reflect.ValueOf(val)

	//3. 取得Kind
	KType := rType.Kind()
	KVal := rVal.Kind()
	fmt.Printf("reflect.Type.Kind = %v  reflect.Value.Kind = %v\n", KType, KVal) //struct struct

	//將reflect.Value 轉成interface{}
	iVal := rVal.Interface()

	fmt.Printf("iVal = %v iVal = %T\n", iVal, iVal) //{austin 10}  main.Student
	//因為我們傳進來參數是interface{}，雖然運行的時候底層知道我們傳進來的是Student struct，
	//但在編譯過程中無法知道，所以取不到stud.Name(透過類型斷言解決)
	//通過類型斷言轉成想要的類型
	stud, ok := iVal.(Student) //或是用switch (assert 章節)
	if ok {
		fmt.Printf("stud.Name = %v\n", stud.Name)
	}
}

//變數、interface{}、reflect.Value互相轉換
func reflectSwift(valInf interface{}) {

	//1. 將interface{}轉為reflect.Value
	reflectVal := reflect.ValueOf(valInf)

	//2.將reflect.Value 轉回interface{}
	intfceVal := reflectVal.Interface()

	//3. 將interface{} 轉成員雷的變數類型 (類型斷言)
	val := intfceVal.(int) //看要轉什麼類型
	fmt.Println(val)
}

//Student struct
type Student struct {
	Name string
	Age  string
}

// 反射

//應用場景 :
//1. json 後面的tag
//2.  適配器函數  (統一處理函數) func(call interface{},args...interface)  =>寫框架常常使用
//3.  底層實例化Action

//基本介紹 :
// 1. 反射可以在運行時，動態獲取變數的各種訊息，如類型(type)、類別(kind)  =>基本數據類型、類別一樣
// 2. 如果是struct 變數，還可以獲取到struct本身的訊息(屬性、方法)
// 3. 反射可以修改變數地值
// 4. 需引入reflect 包

//核心函數
//TypeOf()、ValueOf()   => 用此獲取reflect.Type 、 reflect.Value  類型，使用這兩個類型可以反過來操作變數

//反射常常需要變數、interface{}、reflect.Value互相轉換

//Kind 類別 => 返回一個常數
//Kind跟Type 在基礎數據類型一樣，但如果是 struct  =>  Kind : struct 、 Type : 包名.Student
