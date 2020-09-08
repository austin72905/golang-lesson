package main

import "fmt"

func main() {

	res := sum(10, 20)
	fmt.Println("res=", res) // 4.
	//輸出結果
	// Nodefer res= 32
	// defer2 n2= 20
	// defer1 n1= 10
	// res= 32
}

func sum(n1 int, n2 int) int {

	//當執行到deffer，不會馬上執行，而是先暫存到defer棧中
	defer fmt.Println("defer1 n1=", n1) //3.
	defer fmt.Println("defer2 n2=", n2) //2.
	//defer 放入棧時，也會將相關的值拷貝放入
	//就算下面加，打印出來還是原來的值
	n1++ //n1=11
	n2++ //n2=21

	res := n1 + n2                   //res=32
	fmt.Println("Nodefer res=", res) // 1.
	return res
	//當函數執行完畢，才去把defer 按照先入後出(stack)的方式執行，注意順序!!!!!
}

//defer 延遲機制(有點像await)
//如打開文件、連接資料庫
//作用:在函數執行完，再執行defer 語句，及時釋放資源
//1. 當執行到deffer，不會馬上執行，而是先暫存到defer棧中
//2. 當函數執行完畢，才去把defer 按照先入後出(stack)的方式執行，注意順序!!!!!
//3. defer 放入棧時，也會將相關的值拷貝放入
