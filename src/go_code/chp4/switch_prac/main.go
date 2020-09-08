package main

import "fmt"

func main() {

	//switch 的表達式可以不寫... 類似if else 用法
	var score int = 85
	switch {
	case score >= 90:
		fmt.Println("優秀")
	case score < 90 && score > 60:
		fmt.Println("普通")
	default:
		fmt.Println("不及格")
	}
	//switch 可以直接聲明變數，用分號結尾(不推薦)
	switch tall := 179; {
	case tall > 180:
		fmt.Println("高")
	default:
		fmt.Println("矮")
	}

	//switch 穿透 fallthrough ，如果再case 後面加fallthrough ，則會繼續執行下一個case(不會判斷直接執行)
	var num int = 20
	switch num {
	case 20:
		fmt.Println("20")
		fallthrough //只會穿透一層，要繼續穿透就在下一個case繼續寫
	case 10:
		fmt.Println("10")

	}

	//Type Switch 可以判斷某個interface變數指向的變數類型
	var x interface{}
	var y = 10.0 //float
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的類型: %T", i)
	case float64:
		fmt.Println("x 是 float64型")
	default:
		fmt.Println("未知型")
	}

	//練習
	//根據輸入的月份 打印季節 3~5 春季 6~8夏季  9~11秋  12~2 冬
	var month int16
	fmt.Println("請輸入月分")
	fmt.Scanln(&month)
	switch month {
	case 3, 4, 5:
		fmt.Println("春")
	case 6, 7, 8:
		fmt.Println("夏")
	case 9, 10, 11:
		fmt.Println("秋")
	case 12, 1, 2:
		fmt.Println("冬")
	default:
		fmt.Println("不合規")
	}

}

//switch
//1. 不用加break
//2. 一個case 可以有很多個表達示
//3. 如果每個case都不符合， 執行 default

//細節
//1. case 後是一個表達式(常數、變數、一個有返回值的函數都可以)
//2. case 的表達式數據類型，要和switch 的表達式相同
//3. 一個case 可以有很多個表達示 ，用逗號分開
//4. default 可以不寫
//5. switch 的表達式可以不寫... 類似if else 用法
//6. switch 可以直接聲明變數，用分號結尾(不推薦)
//7. switch 穿透 fallthrough ，如果再case 後面加fallthrough ，則會繼續執行下一個case(不會判斷直接執行)
//8. Type Switch 可以判斷某個interface變數指向的變數類型
