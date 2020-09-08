package tool

import (
	"fmt"
)

//Cal return string
//可以返回多值
//記得寫個英文註釋不然一直警告
func Cal() float64 {
	var num1 float64
	var num2 float64
	var oper string
	var result float64
	//var err string = "輸入格是錯誤"
	fmt.Println("請輸入數字")
	fmt.Scanln(&num1)
	fmt.Println("請輸入數字")
	fmt.Scanln(&num2)
	fmt.Println("請輸入運算元")
	fmt.Scanln(&oper)

	switch oper {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2

	}
	return result
}

//GetSum return int,int
//可返回多個值，如果只想要一個值可用 _ 忽略
//計算 和 和 差
func GetSum(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
