package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//宣告陣列 計算這幾個數的平均值
	hens := [6]float64{3.0, 5.0, 1.0, 3.4, 2.0, 50.0}
	//也可以var hens [6]float64   然後一個一個index給值

	//宣告總和變數
	var totalWei float64
	for i := 0; i < len(hens); i++ {
		totalWei += hens[i]
	}
	//宣告平均值變數
	avrWei := fmt.Sprintf("%.2f", totalWei/float64(len(hens)))

	fmt.Printf("total = %v , avg = %v\n", totalWei, avrWei)

	//宣告之後就算不賦值也有初始值 0
	var intArr [3]int //8字節
	fmt.Println(intArr)
	///地址就是陣列第一個元素的地址
	fmt.Printf("陣列地址 = %p , 陣列第一個元素地址 = %p ,第二個元素地址 = %p\n", &intArr, &intArr[0], &intArr[1])

	//初始化數據的方式
	// var arr1 [3]int = [3]int{1, 2, 3}

	// var arr2 = [3]int{1, 2, 3}
	// var arr3 = [...]int{1, 2, 3} //讓系統判斷有幾個元素
	// var arr4 = [3]int{0: 1, 1: 2, 2: 3} //順序可隨意排 但輸出還是會照順序
	// arr5 := [3]int{1, 2, 3}

	//for-range
	heros := [...]string{"魯夫", "悟飯", "炭治郎"}

	for index, val := range heros {
		fmt.Printf("index = %v , val = %v\n", index, val)
	}

	//陣列 在go裡面是值類型 ， 默認情況下是值傳遞，會進行值拷貝，陣列間不會互相影響
	arr01 := [...]int{1, 2, 3}
	test01(arr01)
	fmt.Println("arr01 =", arr01)
	//如果在其他函數中，要修改原來的陣列，可以用指針
	//把地址傳給指針
	test02(&arr01)
	fmt.Println("arr01 =", arr01)

	//練習
	//創建一個byte類型的陣列，長度26，分別裝A~Z ，把他打印出來   提示:字符數據運算: 'A'+1 ='B'
	var wordArr [26]byte
	var val byte = 'A'
	for i := 0; i < len(wordArr); i++ {

		wordArr[i] = val
		val++
	}
	fmt.Printf("%c\n", wordArr)
	//求出一個陣列的最大值，並得到其所引值
	compareArr := [5]int{5, 7, 1, 10, 3}
	//假設地一個元素是最大的
	maxVal := compareArr[0]
	maxValIndex := 0
	//讓每個元素旱地一個元素逐一比較
	for i := 1; i < len(compareArr); i++ {
		//如果arr[i]值比第一個大，變化maxVal 的值
		if compareArr[i] > maxVal {
			maxVal = compareArr[i]
			//只有arr[i]值比第一個大才會變
			maxValIndex = i
		}
	}
	fmt.Printf("最大值 = %v , 索引值 = %v\n", maxVal, maxValIndex)
	//求出一個數組的和與平均值 for-range
	var totalComArr int
	var avrComArr float64
	for _, val := range compareArr {
		totalComArr += val
	}
	avrComArr = float64(totalComArr) / float64(len(compareArr))
	fmt.Printf("compareArr總和 = %v , 平均 = %v\n", totalComArr, avrComArr)

	//隨機生成5個數，並反轉打印
	//反轉打印  交換次數是 len/2  倒數第n個元素，與n個交換
	rand.Seed(time.Now().UnixNano())
	var ranArr [5]int
	ranArrIdx := len(ranArr) - 1
	chgTime := len(ranArr) / 2

	for i := 0; i < len(ranArr); i++ {
		ranArr[i] = rand.Intn(100) //0<=n<100
	}
	fmt.Println("交換前", ranArr)

	temp := 0
	//len(ranArr)/2 執行交換的次數
	for i := 0; i < chgTime; i++ {
		temp = ranArr[ranArrIdx-i]
		//倒數第n個元素，與n個交換
		ranArr[ranArrIdx-i] = ranArr[i]
		ranArr[i] = temp
	}
	fmt.Println("交換後", ranArr)
}

//這樣陣列是值拷貝到函數理
func test01(arr [3]int) {
	arr[0] = 88
}

//拿到了陣列的地址
func test02(arr *[3]int) {
	//先取到指向的陣列
	(*arr)[0] = 88
}

//陣列

//陣列 在go裡面是值類型 ， 默認情況下是值傳遞，會進行值拷貝，陣列間不會互相影響

//占用連續記憶體空間

//宣告之後就算不賦值也有初始值
//1. 數字 默認值: 0
//2. string 默認值: ""
//3. bool 默認值: false

//地址就是陣列第一個元素的地址

//下一個元素地址 是第一個的地址 +  數據類型佔的字節數

//for-range  (for index,value := range arr{....})

//1. index 陣列索引值
//2  value 各元素的值
//3. 變數作用域在for裡面
//4. index、value想忽略可用_
//5. index、value 名稱可自行指定

//如果沒有寫陣列長度的話 var arr []int  是宣告成slice 不是陣列

//如果out of range 會報 panic(恐慌) 錯誤

//如果在其他函數中，要修改原來的陣列，可以用指針

//長度也是陣列類型的一部份，傳遞函數參數時，要考慮陣列的長度  func test(arr [4]int)   不能帶arr [3]int 進去
