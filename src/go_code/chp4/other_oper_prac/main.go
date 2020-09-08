package main

import "fmt"

func main() {

	// & 和* 使用
	//&a a變數的地址
	a := 100
	fmt.Println("a 的地址", &a)

	// * 代表是指針變數  * =>可以理解為指向某個地址(所以叫指針變數)
	var ptr *int = &a
	fmt.Println("ptr 指向的值是", *ptr)

	//&  : 兩位都為1 , 結果為 1 否則為 0
	//|  : 兩位有一個1 , 結果為 1 否則為 0
	//^  : 兩位有一個1、一個0 , 結果為 1 否則為 0
	//ex : 2&3 =? 等於2
	//  2 補碼[0000 0010]
	//  3 補碼[0000 0011]
	//=> &    [0000 0010] => 2
	fmt.Println(2 & 3)
	//ex : 2|3 =?
	//=> |    [0000 0011] => 3
	fmt.Println(2 | 3)
	//ex : 2^3 =?
	//=> ^    [0000 0001] => 1
	fmt.Println(2 ^ 3)
	//ex : -2^2=?
	// -2 原碼[1000 0010] 反碼[1111 1101] 補碼[1111 1110]
	// 2                                 補碼[0000 0010]
	// => ^  [1111 1100] (這個處理完還是補碼...)
	// =>反碼[1111 1011] =>原碼[1000 0100] => -4
	fmt.Println(-2 ^ 2)

	//<< : 溢出低位, 符號位不變，用符號位補益出的高位
	//ex c:=1>>2(右移兩位)  0000 0001 => 0000 0000 =>0
	c := 1 >> 2
	fmt.Println(c)
	//>> : 符號位不變, 低位補0
	//ex d:=1<<2(左移兩位)  0000 0001 => 0000 0100 =>4
	d := 1 << 2
	fmt.Println(d)

}

//先看後面的進制
//其他運算符
//位移運算符
//&  : 兩位都為1 , 結果為 1 否則為 0
//|  : 兩位有一個1 , 結果為 1 否則為 0
//^  : 兩位有一個1、一個0 , 結果為 1 否則為 0
//<< : 溢出低位, 符號位不變，用符號位補益出的高位
//>> : 符號位不變, 低位補0

// & 和* 使用
//沒有三元運算符