package main

import "fmt"

func main() {

	monkey := littleMonkey{
		monkey{
			name: "悟空",
		},
	}

	monkey.climbing()
	monkey.flying()
	monkey.swimming()
}

type monkey struct {
	name string
}

func (mon *monkey) climbing() {
	fmt.Println(mon.name, "會爬樹")
}

type littleMonkey struct {
	monkey //繼承
	//birdCanDo //在實作的時候不用把他加到屬性
}

//實作birdCanDo介面

func (litMin *littleMonkey) flying() {
	fmt.Println(litMin.name, "學會了飛翔")
}

func (litMin *littleMonkey) swimming() {
	fmt.Println(litMin.name, "學會了游泳")
}

//介面
type birdCanDo interface {
	flying()
}

//介面
type fishCanDo interface {
	swimming()
}

//繼承與介面

//當a struct 繼承了b struct，a 自動繼承了b 的屬性和方法，可以直接使用

//不破壞繼承關係的情況下，進行功能的擴展

//繼承: 解決代碼重複

//介面: 設計各種規範， 讓其他自訂義類型去實作
