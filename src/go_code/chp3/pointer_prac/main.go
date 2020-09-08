package main

import "fmt"

func main() {
	//基本數據的內存
	var i int = 10
	// i的地址 用&
	fmt.Println("i 的地址:", &i)
	//指針變數 用* =>也會占用空間(本身也有地址)，存的是指向的地址
	//類型 *int
	//本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr:%v\n", ptr)
	fmt.Printf("ptr 的地址:%v\n", &ptr)
	//指針變數取到指向位置的值 用*
	fmt.Printf("ptr 指向的值:%v\n", *ptr)

	//練習
	//1. 獲取一個int變數 num 的地址，顯示到終端機
	//2. 給num的地址給予指針 pitr,並通過pitr 修改 num 的值
	//1.
	var num int = 10
	fmt.Printf("num 的值:%v\n", num)
	fmt.Printf("num 的地址:%v\n", &num)
	//2.
	var pitr *int = &num
	*pitr = 9
	fmt.Printf("num 的值:%v\n", num)
}

//指針
//1. 基本數據的內存
//2. 指針變數
//3. 指針變數取到指向位置的值
//4. 值類型都有自己的指針類型 =>*數據類型
//5. 值類型 : int float bool string 數組 strct
