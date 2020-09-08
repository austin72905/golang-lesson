package main

import "fmt"

func main() {

	str := "helloWorld" //string 本身也是切片，地址指向byte陣列的第一個元素
	strSlice := str[5:]
	fmt.Println("strSlice =", strSlice)
	fmt.Printf("strSlice 類型 =%T\n", strSlice) //string

	//修改字串方法
	//這種方法是用字節，中文佔3字節，會亂碼
	byteArr := []byte(str)
	//修改
	byteArr[0] = 's'
	//轉回來
	str2 := string(byteArr)
	fmt.Println("str2 = ", str2)

	//修改中文
	//rune 是用字符處理
	str3 := "舔中共屁眼"
	runeArr := []rune(str3)

	runeArr[0] = '吸'

	str4 := string(runeArr)
	fmt.Println("str4 = ", str4)
}

//string 和 slice 關係

//1. string 底層是byte陣列 ，因此可以用slice處理

//2. string 本身也是切片，地址指向byte陣列的第一個元素

//3. sting 不可變，不能用str[0]='a'，來修改字符串

//4. 如果要改變，要先將str 轉成[]byte / []rune  ，修改完再轉回字串
