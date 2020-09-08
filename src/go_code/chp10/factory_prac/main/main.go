package main

import (
	"fmt"
	"go_code/chp10/factory_prac/model"
)

func main() {

	//不給偷懶欸...
	//如果要引用的是大寫OK
	var stu = model.Student{
		Name:  "austin",
		Score: 64.1,
	}
	fmt.Println(stu)
	//如果是小寫
	//工廠模式
	var stu2 = model.NewStu2("TO", 63.1)              //這是一個指針變數
	fmt.Println("名字=", stu2.Name, "分數 =", stu2.Score) //標準要(*stu2). 才能取到值

	//如果裡面有個屬性是小寫...，要例外寫個大寫方法去拿(也叫做封裝)
	var stu3 = model.NewStu3("sam", 65.0)
	fmt.Println("名字=", stu3.Name, "分數 =", stu3.GetScore())
}

//工廠模式

//go 裡面沒有建構子，如果有一個model 包的Student 結構體，想要創建他的實例，就直接引用，但如果是小寫就無法引用了，所以需要工廠模式
