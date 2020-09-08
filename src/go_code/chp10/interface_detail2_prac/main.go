package main

import "fmt"

func main() {

	var stu student = student{}

	//會報錯，因為student 是定義一個指針方法
	//var per person = stu //student does not implement person (say method has pointer receiver)

	//要修改成這樣
	var per person = &stu
	per.say()
}

//測試一

//Ainter is to struct with Title in testing
type Ainter interface {
	Test1()
}

//Binter is to struct with Title in testing
type Binter interface {
	Test1()
}

//Cinter is to struct with Title in testing
type Cinter interface {
	Ainter
	Binter
}

//測試二

type person interface {
	say()
}

type student struct {
}

//定義一個student 的"指針"方法
func (stu *student) say() {
	fmt.Println("test...")
}

//影片上說 介面繼承介面，繼承的介面不能有相同方法，但我測試沒有
