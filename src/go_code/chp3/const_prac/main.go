package main

import "fmt"

func main() {
	const test = 1
	fmt.Println(test)

	//定義一個變數
	var varN int
	varN = 4
	//const test2 = varN / 1
	//這樣會報錯，因為varN 是個變數
	fmt.Println(varN)

	//還有一個專業寫法 (我覺得就是C#的enum)
	//這個寫法第一個值=0 依序類推(一行遞增一次)
	//一樣可以用大小寫控制其他包是否可以使用

	const (
		alipay = iota
		wexin
		fastpay, unionpay = iota, iota
		qqpay             = iota
	)
	//一般都只在第一個常數寫iota
	//0 1 2 2 3
	fmt.Println(alipay, wexin, fastpay, unionpay, qqpay)
}

//常數

//1. 要用const 聲明
//2. 在被定義時必須初始化 (必須給他值)
//3. 不能被修改
//4. 只能修飾bool、數值類、string
