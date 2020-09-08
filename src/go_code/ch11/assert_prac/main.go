package main

import "fmt"

func main() {
	//定義一個空介面
	var a interface{}

	var p1 point = point{1, 2}

	a = p1 //ok  因為a 是空介面
	//如何將a 再賦予一個point變數?
	var p2 point
	p2 = a.(point) //類型斷言 如果a 可以轉成 point，就可以賦值給p2   (a 也要先指向point 才可以轉)

	fmt.Println(p2)
	fmt.Printf("p2 的類型是 %T 值 = %v\n", p2, p2)

	//也可以帶檢查 有點像 tryParse
	var x interface{}
	var num1 float32 = 1.3
	x = num1               //OK  空接口可接收任何類型
	y, isOk := x.(float64) // isOk 是一個bool

	if isOk {
		fmt.Println("轉換成功")
		fmt.Printf("y 的類型是 %T 值 = %v\n", y, y)
	} else {
		fmt.Println("轉換失敗")
	}

	fmt.Println("運用檢查機制，就算失敗可以繼續執行...")
}

type point struct {
	x int
	y int
}

//類型斷言
