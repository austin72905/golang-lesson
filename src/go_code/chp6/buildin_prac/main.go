package main

import "fmt"

func main() {
	//一般的值
	num1 := 100
	fmt.Printf("num1 的 類型 = %T,  值 = %v ,  地址 = %v\n", num1, num1, &num1)

	//new 出來的指針
	//num2 的 類型 => *int
	//num2 的 值 => 0xc00000c178  (每次都不一樣，系統看情況分配的)
	//num2 的 地址 => 0xc000006030 (每次都不一樣，系統看情況分配的)
	//num2 的 指向地址的值 => 0
	num2 := new(int) //*int  指向地址所擁有的值int

	fmt.Printf("num2 的 類型 = %T,  值 = %v ,  地址 = %v , 指向的地址擁有的值為 = %v\n", num2, num2, &num2, *num2)

	//想要修改值
	*num2 = 10
	fmt.Printf("num2 的 類型 = %T,  值 = %v ,  地址 = %v , 指向的地址擁有的值為 = %v\n", num2, num2, &num2, *num2)
}

//內建函數介紹

//new   (基本類型)系統分配內存 ，並返回一個指針，該指針指向的地址的值為0  func new(Type) *Type

//make  (引用類型)
