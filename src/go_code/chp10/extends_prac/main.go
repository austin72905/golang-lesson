package main

import "fmt"

func main() {

	//呼叫小學生
	per := &Child{}
	//per.Student.Name = "austin"
	//per.Student.Age = 8
	//匿名結構體可以簡化，他會先看Child 裡有沒有這個屬性，沒有就去找匿名結構體Studnet裡有沒有
	//如果strcut 跟 匿名結構體都有 一樣的屬性，賦值就要明確了
	per.Name = "austin"
	per.Age = 8
	per.testing()
	per.SetScore(50)
	per.ShowInfo() // 名字會印出空字串，因為上面只有宣告到 Child裡面的Name
	//per.Student.Name = "austin"
	//per.Name = "austin"    這兩個是不同的東西
	per.Student.Name = "austin"
	per.ShowInfo()

	//呼叫大學生
	per1 := &College{}
	per1.Name = "mary"
	per1.Age = 18
	per1.testing()
	per1.SetScore(73)
	per1.ShowInfo()
}

//ex: 大學生、小學生

//Student struct
//將共有屬性提取出來
type Student struct {
	Name  string
	Age   int
	Score int
}

//將共有方法綁定 Student

//ShowInfo ...
func (stud *Student) ShowInfo() {
	fmt.Printf("學生名=%v 年紀=%v 成績=%v\n", stud.Name, stud.Age, stud.Score)
}

//SetScore ...
func (stud *Student) SetScore(score int) {
	stud.Score = score
}

//小學生struct，繼承Student

//Child struct
type Child struct {
	Student        //繼承Student 匿名結構體
	Name    string //如果strcut 跟 匿名結構體都有 一樣的屬性
}

//小學生特有方法
func (child *Child) testing() {
	fmt.Println("小學生test")
}

//大學生

//College struct
type College struct {
	Student //繼承Student 匿名結構體
}

//大學生特有方法
func (college *College) testing() {
	fmt.Println("大學生test")
}

//繼承

//宣告方法: 使用匿名結構體來繼承，把一個struct 寫在另一個struct裡面 (嵌套以後可以使用該struct 所有的屬性和方法，不論大小寫)

//如果strcut 跟 匿名結構體都有 一樣的屬性或方法， 會採就近方法

//可以同個struct 繼承 多個 匿名結構體
