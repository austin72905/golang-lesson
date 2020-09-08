package main

import "fmt"

//全局匿名函數，記得大寫
var (
	//不是:=
	Fun1 = func(n1 int, n2 int) int {
		return n1 + n2
	}
)

func main() {
	//要使用直接用，後面直接傳參數
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 6)
	fmt.Println("res1=", res1)

	//也可將函數存成函數變數，可以調用a
	a := func(n1 int, n2 int) int {
		return n1 + n2
	}
	res2 := a(10, 2)
	fmt.Println("res2=", res2)

	//調用全局匿名函數
	res3 := Fun1(5, 3)
	fmt.Println("res3=", res3)
}

//匿名函數
//要使用直接用，後面直接傳參數
//也可將函數存成函數變數，可以調用a
//全局匿名函數
