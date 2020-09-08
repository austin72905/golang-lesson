package main

import "fmt"

func main() {
	var a byte = 'a'
	var b byte = '0'
	fmt.Println(a, b)            //印出的是'a','0'的ascii碼的值
	fmt.Printf("%c  %c\n", a, b) //如果希望印出對應的字符%c，要使用格式化輸出
	//漢字
	var be int = '北'
	fmt.Printf("%c %d\n", be, be) //%d 印出值

	var c int = 22269
	fmt.Printf("%c\n", c) //會印出北
}

//字符 char
//1. go 沒有專門的字符類型，都用byte
//2. go 的string 不是由char組成，而是由字節組成
//3. go 語言是用 UTF-8
//4. 英文 : 1字節、漢字 : 3字節
//5. 字符本質是一個整數
//6. 字符是一個整數，所以可以運算
