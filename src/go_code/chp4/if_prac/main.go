package main

import (
	"fmt"
	"math"
)

func main() {

	//golang支持在if語句宣告變數
	//一定要大括號
	if age := 20; age > 18 {
		fmt.Println("滿18了")
	} else {
		fmt.Println("未滿")
	}

	//練習
	//1. 聲明兩個int 變數  若兩數之和>50 打印hello world
	var num1 int32
	var num2 int32
	fmt.Println("請數入兩個數字")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	if num1+num2 > 50 {
		fmt.Println("hello world")
	}
	//2. 聲明兩個float64 變數 第一個數>10.0 且 第二個數<20.0，打印總和
	var num3 float64
	var num4 float64
	fmt.Println("請數入兩個數字")
	fmt.Scanln(&num3)
	fmt.Scanln(&num4)
	if num3 > 10.0 && num4 < 20 {
		fmt.Println(num3 + num4)
	}
	//3. 定義兩個變數 int32 判斷其總和 ，能否被3又被5整除 ，如果可以打印OK
	var num5 int32
	var num6 int32
	fmt.Println("請數入兩個數字")
	fmt.Scanln(&num5)
	fmt.Scanln(&num6)
	if (num5+num6)%3 == 0 && (num5+num6)%5 == 0 {
		fmt.Println("OK")
	}
	//4. 判斷某年是否是閏年   兩者之1  1) 年份能被4整除，但不能被100整除 2) 能被400整除
	fmt.Println("請數入年分")
	var year int
	fmt.Scanln(&year)
	// if year%400 == 0 {
	// 	fmt.Println("是閏年")
	// } else if year%4 == 0 {
	// 	if year%100 != 0 {
	// 		fmt.Println("是閏年")
	// 	}
	// }
	//y最佳解
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("是閏年")
	}

	//5. 小名參加考試，與爸爸約定
	//如果成績100分， 爸爸送BMW
	//如果成績80~99 ，送IPONE
	//如果60~80 送IPAD
	//其他什麼都沒有
	fmt.Println("小明考幾分?")
	var score int
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("得到BMW")
	} else if score >= 80 {
		fmt.Println("得到IPONE")
	} else if score >= 60 {
		fmt.Println("得到IPAD")
	} else {
		fmt.Println("小明一無所有")
	}

	//求ax^2+bx+c 的根，a,b,c分別回函數的參數 ， 如果 b^2-4ac>0 則有兩個解
	//b^2-4ac=0  ,則有一個解  ,b^2-4ac<0 ,則無解

	//提示1   x1=(-b+sqrt(b2-4ac))/2a
	//       x2=(-b-sqrt(b2-4ac))/2a
	//提示2  math.Sqrt(num);  可以求平方根
	//測試數據 3,100,6

	var a float64 = 3.0
	var b float64 = 100.0
	var c float64 = 6.0

	m := b*b - 4*a*c
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b - math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v  x2=%v\n", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v\n", x1)
	} else {
		fmt.Println("無解")
	}

	// 嫁女兒 條件:　身高：180cm  財富: 1千萬以上  帥 : 是
	//1. 如果三個條件都滿足 打印 "我一定嫁"
	//2. 至少有一個滿足  打印  "勉強"
	//3. 都不滿足  打印  "不嫁"
	var tall float64
	var money int32
	var isGood bool
	fmt.Println("請輸入擇偶條件")
	fmt.Scanln(&tall)
	fmt.Scanln(&money)
	fmt.Scanln(&isGood) //true or false
	if tall > 180.0 && money > 10000000 && isGood {
		fmt.Println("我一定嫁")
	} else if tall > 180.0 || money > 10000000 || isGood {
		fmt.Println("勉強")
	} else {
		fmt.Println("不嫁")
	}

}
