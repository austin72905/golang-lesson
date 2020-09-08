package main

import "fmt"

func main() {

	for i := 0; i < 4; i++ {

		for j := 0; j < 5; j++ {

			if j == 2 {
				continue
			}
			fmt.Println("j=", j) //每次都沒有2
		}
	}
	//也可以加入標籤
label1:
	for i := 0; i < 4; i++ {

		for j := 0; j < 5; j++ {

			if j == 2 {
				continue label1 //01 01 01 01 01 每次到2都會跳過跑到"標籤所在的迴圈"的下一個迭代
			}
			fmt.Println("j=", j)
		}
	}

	//goto  可以無條件轉移到指定的行，不太推薦使用
	// 	fmt.Println("hello")

	// 	goto label2
	// 	//以下都不會執行
	// 	fmt.Println("hello")
	// 	fmt.Println("hello")
	// 	fmt.Println("hello")
	// 	//會跳到這
	// label2:
	// 	fmt.Println("world")
	// 	fmt.Println("world")
	// 	fmt.Println("world")

	//return 是直接跳出函數
}

//continue 結束本次迴圈，執行下一次迴圈

//也可以用標籤表示要跳過哪個迴圈

//不再執行循環體，直接到循環迭代
