package main

import "fmt"

func main() {
	var a string = "hello"
	fmt.Println(a)
	//使用反引號 ，可把裡面的內容原汁原味呈現
	b := `abv\nabc`
	fmt.Println(b)
	//可用 + 拼接
	c := "hello" + "world"
	fmt.Println(c)
	//如果很長要換行  + 號要留在上面
	d := "aa" + "bb" +
		"cc"
	fmt.Println(d)

	//默認值
	var e int     //0
	var f float32 //0
	var g float64 //0
	var h bool    //false
	var i string  //""
	//格式化輸出
	//%v : 按變數原始值輸出
	fmt.Printf("%d %f %f %v %v", e, f, g, h, i)

}

//1. go 的string 不是由char組成，而是由字節組成
//2. string 字符的內容不能換
//3. 可使用反引號
//4. 可用 + 拼接
