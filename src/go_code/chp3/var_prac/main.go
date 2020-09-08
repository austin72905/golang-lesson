package main

import "fmt"

//全域變數聲明方式
var (
	n7 = 7
	n8 = 8
	n9 = 9
)

func main() {
	//(1)聲明後不給值，會使用預設值 int =>0
	var i int
	fmt.Println(i)
	//(2)根據值，自行判斷類型
	var num = 10
	fmt.Println(num)
	//(3)省略var  :=   等於 var name string="ausitn"
	name := "ausitn"
	fmt.Println(name)
	//4. 多變數聲明
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)
	//多變數聲明 (類型不同)
	var n4, n5, n6 = 10, "lin", 3 //或  n4, n5, n6 := 10, "lin", 3
	fmt.Println(n4, n5, n6)

	fmt.Println(n7, n8, n9)
}

//注意事項:
// 1. 變數在宣告時就會在內存佔據一個空位
// 2. 該空位有自己的名稱和 數據類型(int)
// 3. (1)聲明後不給值，會使用預設值
//    (2)根據值，自行判斷類型
//    (3)省略var
// 4. 多變數聲明
// 5. 全局變數聲明

// Go的基本類型有basic types

// bool

// string  ==>字串是基本類型

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // uint8 的別名

// rune // int32 的別名  ==>字符串裡有中文都用它存
//      // 代表一個Unicode碼

// float32 float64

// complex64 complex128
