package main

import "fmt"

func main() {

	//這樣寫會報錯，介面本身不能創建實例
	// var p person
	// person.say()
	var stu student
	//可以指向實現該介面的變數
	var p person = stu
	p.say()

	//只要是自訂義類型，都可以實現接口，不一定要struct
	var n myName = "austin"
	var b person = n
	b.say()

	// 一個介面(A)可以繼承多個介面(B、C)，如果要實作A介面，(B、C)介面的方法也要全部實作
	var wolongNoodle woLongNoodle
	//wolongNoodle 有實作addNoodle()、addSomeWater()、addSomeSauce()
	var cookNoodle cook = wolongNoodle
	cookNoodle.addNoodle()
	cookNoodle.addSomeWater()
	cookNoodle.addSomeSauce()
	cookNoodle.test()

	//interface{} 空介面，所有類型都實現了空接口，所以我們可以把任何一個變數給空介面
	var emp empty = wolongNoodle
	fmt.Println(emp)

	//或是直接在這宣告空介面
	var emp2 interface{} = n
	fmt.Println(emp2)
	num := 6.4
	emp2 = num
	fmt.Println(emp2)
}

type person interface {
	say()
}

type student struct {
	name string
}

func (stu student) say() {
	fmt.Println("student say...")
}

//只要是自訂義類型，都可以實現接口，不一定要struct
type myName string

func (n myName) say() {
	fmt.Println("my name is", n)
}

// 一個介面(A)可以繼承多個介面(B、C)，如果要實作A介面，(B、C)介面的方法也要全部實作
//ex  煮麵 : 要加水、加醬料包
type cook interface {
	addNoodle()
	//繼承另外兩個介面(影片上說介面繼承介面，繼承的介面不能有相同方法，但我測試沒有)
	addWater
	addSauce
}

type addWater interface {
	test()
	addSomeWater()
}

type addSauce interface {
	test()
	addSomeSauce()
}

type woLongNoodle struct {
}

//實作
func (w woLongNoodle) addNoodle() {
	fmt.Println("今天加烏龍麵...")
}

func (w woLongNoodle) addSomeWater() {
	fmt.Println("加85。c的熱水...")
}

func (w woLongNoodle) addSomeSauce() {
	fmt.Println("加點味噌...")
}

func (w woLongNoodle) test() {
	fmt.Println("test...")
}

//interface{} 空介面
type empty interface {
}

//接口(介面)細節

//1. 不能創建實例，但可以指向實現該介面的變數

//2. 介面沒有方法體，即沒有實現的方法

//3. 實現該介面，是指一定要實作介面的所有方法

//4. 只要是自訂義類型，都可以實現介面，不一定要struct

//5. 一個自訂義類型可以實現多個介面(該自訂義類型所擁有的方法，只要包含某個介面所有的方法，就算是實現該介面了)

//6. 介面裡面不可以有任何的變數

//7. 一個介面(A)可以繼承多個介面(B、C)，如果要實作A介面，(B、C)介面的方法也要全部實作

//8. interface類型默認是指針(引用類型)，如果沒有對interface初始化就使用，會輸出nil

//9. interface{} 空介面，所有類型都實現了空接口，所以我們可以把任何一個變數給空介面
