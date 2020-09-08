package main

import (
	"fmt"
)

func main() {
	//break練習
	//1. 寫一個無限迴圈，不停生成隨機數(1~100)，當生成99個，break

	//隨機數 math/rand 包 func Intn(n int) int ，反回一个取值范围在[0,n)的伪随机int值，如果n<=0会panic
	//但是要設定一個seed不然生成的隨機數會固定...
	//func().Unix() =>會生成一個1970/1/1 00:00:00~前面涵式指定時間的秒數
	//rand.Seed(time.Now().Unix()) //他會返回一個秒數 1970/1/1 00:00:00 ~~~現在的秒數
	//time.Now().Unix()永遠都不一樣，所以Seed 會永遠不一樣，就可以創造隨機數
	//s := rand.Intn(100) + 1 //但不包含100，所以+1
	//fmt.Println(s)

	// var count int = 0
	// for {
	// 	rand.Seed(time.Now().UnixNano()) //Unix()太慢所以改單位
	// 	n := rand.Intn(100) + 1
	// 	count++
	// 	if n == 99 {
	// 		break
	// 	}
	// }
	// fmt.Println("得到99 共花了", count, "次")

	//默認跳出最近的break
	// for i := 0; i < 4; i++ {
	// 	for j := 0; j < 10; j++ {

	// 		if j == 2 {
	// 			break //默認跳出最近的break
	// 		}
	// 		fmt.Println("j=", j) //01 01 01 01
	// 	}
	// }

	//可以指定標籤來決定要break哪段
label2:
	for i := 0; i < 4; i++ {
		//label1: //標籤，名稱可自訂
		for j := 0; j < 10; j++ {

			if j == 2 {
				//break label1 //break 標籤的for迴圈
				break label2
			}
			fmt.Println("j=", j) //01 01 01 01
		}
	}

	//練習
	//1. 100內的數求和。當和 第一次>20 的當前數
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
		if sum >= 20 {
			fmt.Println("加到", i, "會大於20")
			break
		}
	}
	//2. 登入有3次機會，如果戶名為:austin  密碼:000  打印登入成功 ，否則提示還有幾次機會
	var account string
	var psrd string
	var chance int = 3
	for i := 1; i <= chance; i++ {
		fmt.Println("請輸入帳號")
		fmt.Scanln(&account)
		fmt.Println("請輸入密碼")
		fmt.Scanln(&psrd)

		if account == "austin" && psrd == "000" {
			fmt.Println("登入成功")
			break //跳出最近的for迴圈
		}

		//判斷機會還有幾次
		if chance-i == 0 {
			fmt.Println("登入失敗")
			break
		}
		fmt.Println("還有", chance-i, "次機會")
	}

}

//1. 可以指定標籤來決定要break哪段
//2. 預設break最近的for迴圈
