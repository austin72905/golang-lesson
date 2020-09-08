package main

import "fmt"

func main() {
	fmt.Println("請輸入姓名")
	var name string
	var score int
	// 他就直接把輸入值存入了該變數的地址
	fmt.Scanln(&name)
	fmt.Println("你輸入了", name)
	fmt.Println("---Scanf練習---")
	fmt.Println("請輸入姓名、分數，以空格隔開")
	fmt.Scanf("%v %v", &name, &score)
	fmt.Println("你輸入了", name, "分數", score)
}

//鍵盤輸入
//Scanln =>掃描該行
//Scanf =>可自訂輸入格式
