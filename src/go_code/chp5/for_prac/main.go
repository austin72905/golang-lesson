package main

import "fmt"

func main() {

	//可以這樣寫，但不能用var
	for i := 0; i < 10; i++ {
		fmt.Println("你好")
	}

	//第二種寫法
	j := 0       //循環變數初始化
	for j < 10 { //循環條件
		fmt.Println("很好") //循環體
		j++               //循環變數迭代
	}

	//第三種寫法
	//無限循環，通常配合break
	k := 0
	for {
		fmt.Println("安安")
		k++
		if k > 10 {
			break
		}
	}

	//for-range  歷遍 集合的方法
	//1. 傳統
	//但這種方法，如果str有中文 會亂碼，因為這個歷遍是用字節，但中文在utf-8 對應3個字節
	//解決: 將str 轉成 []rune 切片
	var str string = "hello台灣"
	str2 := []rune(str)
	//len()長度
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}

	//2. for-range
	for index, val := range str {
		fmt.Printf("index=%d, val=%c\n", index, val)
		//會打印出
		//注意看index
		//index=0, val=h
		//index=1, val=e
		//index=2, val=l
		//index=3, val=l
		//index=4, val=o
		//index=5, val=台
		//index=8, val=灣
	}

	//練習
	//1. 打印 1~100 所有是9的倍數 有幾個 和總和
	var count int //用uint 也是可
	var total int
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			total += i
			count++
		}
	}
	fmt.Printf("9的倍數有%v個 , 總和為%v", count, total)
}

//for 執行順序
//1. 循環變數初始化
//2. 循環條件 =>判斷是否true ，不是就結束 他是一個返回bool的判斷式
//3. 循環體
//4. 循環變數迭代
