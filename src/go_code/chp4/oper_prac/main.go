package main

import "fmt"

func main() {
	//1. 如果運算都是整數，結果只會是整數
	fmt.Println(10 / 4)

	//如果希望有小數，要有小數進行運算
	var a float32 = 10.0 / 4
	fmt.Println(a)

	//2. ++ 只能獨立使用 ex: a=i++(會報錯) i=i++ (也會錯) 要先i++ 再把i 賦予a
	//以下會報錯
	// var b int =0
	// if b++>0{

	// }

	//3. ++只能在變數後面 ++a 會報錯

	//練習題
	//1. 假設有97天放假，那是x個星期又y天?
	var x int = 97 / 7
	var y int = 97 % 7
	fmt.Printf("%d星期又%d天\n", x, y)
	//fmt.Println(x, "星期", y, "天")

	//2. temp變數表示華氏溫度 請求出其攝氏溫度
	var wtemp float32 = 12.6
	var ctemp float32 = (wtemp - 100) * 5.0 / 9 //如果寫5會除出整數
	fmt.Println(ctemp)

	// && 又稱短路與   第一個為false 後面就不判斷了 => 結果false
	// || 又稱短路或   第一個為true  後面就不判斷了 => true

	//賦值運算符 += *= ...

	//練習
	//有兩個變數big,small，要求將其進行交換，但不允許是用中間變數
	//用中間變數的解法
	fmt.Println("---------練習---------")
	// big := 10
	// small := 1
	// fmt.Println("big的值=", big, "small的值=", small)
	// //中間變數
	// temp := big
	// big = small
	// small = temp
	// fmt.Println("big的值=", big, "small的值=", small)

	//不能用中間數的解法
	big := 10
	small := 1
	fmt.Println("big的值=", big, "small的值=", small)
	big = big + small   // big 11
	small = big - small // small =11 -1 =10
	big = big - small   //big 11-10=1
	fmt.Println("big的值=", big, "small的值=", small)

}
