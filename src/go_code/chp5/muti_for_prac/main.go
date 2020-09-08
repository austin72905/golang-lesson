package main

import "fmt"

func main() {
	//練習
	//1. 有3個班級，每班有5個人
	//    (1)求各班平均分  與 所有班的平均分
	//    (2) 統計及格人數
	// var sum float32
	// var totalSum float32
	// var student = 5
	// var classNum = 3
	// var match int
	// for k := 1; k <= classNum; k++ {
	// 	for i := 1; i <= student; i++ {
	// 		var score float32
	// 		fmt.Println("請輸入成績")
	// 		fmt.Scanln(&score)
	// 		//判斷是否及格
	// 		if score >= 60 {
	// 			match++
	// 		}
	// 		//計算總分 才可以計算平均
	// 		sum += score
	// 	}
	// 	fmt.Printf("第%v班 平均%v\n", k, sum/float32(student)) //用變數缺點就是型態不一樣他會報錯，所以還要型轉
	// 	totalSum += sum
	// }
	// fmt.Printf("總平均分數為%v\n", totalSum/float32(student*classNum))
	// fmt.Println("即個人數為", match)

	//金字塔
	//遇到很複雜的先把問題縮小，一個一個解決他
	//找規律

	//第一版
	//*
	//**
	//***
	//第幾層就印幾個*
	// for i := 1; i <= 3; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	//第二板
	//每層*數  2* 層數 -1
	//*
	//***
	//*****
	// for i := 1; i <= 3; i++ {

	// 	for j := 1; j <= 2*i-1; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
	//第三版
	//加入空格
	//從左邊數 要空幾格

	//金字塔   每層*數                    空格
	//  *     2* 層數 -1     1      總層數-當前層數           2
	// ***    2* 層數 -1     3                               1
	//*****   2* 層數 -1     5                               0
	for i := 1; i <= 3; i++ {
		//在打印* 要先打印空格
		//3-i (總層數-當前層數)
		for k := 1; k <= 3-i; k++ {
			fmt.Print(" ")
		}

		//打印 *
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//空心金字塔
	//跟前面一樣，只是每層在空格完之後，只有 "第一個"(1) 跟 "最後一個"(2*i-1)打印* ，其他都打印空格
	for i := 1; i <= 3; i++ {
		//在打印* 要先打印空格
		//3-i (總層數-當前層數)
		for k := 1; k <= 3-i; k++ {
			fmt.Print(" ")
		}

		//打印 *
		for j := 1; j <= 2*i-1; j++ {
			//每層在空格完之後，只有 "第一個"(1) 跟 "最後一個"(2*i-1)打印* ，其他都打印空格
			//如果是最後一行也是全部打*
			if j == 1 || j == 2*i-1 || i == 3 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

}
