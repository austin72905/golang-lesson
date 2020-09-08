package main

import "fmt"

func main() {

	//1. 順序查找
	//練習
	//有一陣列 白,黃,藍,紅，鍵盤輸入，判斷陣列是否有輸入的顏色
	color := [4]string{"白", "黃", "藍", "紅"}
	var input string
	fmt.Println("請輸入顏色")
	fmt.Scanln(&input)

	index := -1
	for i := 0; i < len(color); i++ {
		if input == color[i] {
			index = i
		}
	}
	if index != -1 {
		fmt.Printf("你輸入的顏色在%v位\n", index)
	} else {
		fmt.Println("你輸入的顏色不再陣列裡")
	}
	//二分查找 (一直折半，很像終極密碼)
	//練習

	arr1 := [6]int{1, 8, 10, 89, 1000, 1234}

	binaryFind(&arr1, 0, len(arr1)-1, 8)

}

func binaryFind(arr *[6]int, leftIdx int, rightIdx int, val int) {

	//退出遞迴條件 =>左右標會不斷靠近，當 左>右 還找不到就代表 真的找不到了
	if leftIdx > rightIdx {
		fmt.Println("找不到")
	}

	midIdx := (leftIdx + rightIdx) / 2
	//沒找到就繼續找阿，所以用遞迴
	if (*arr)[midIdx] > val {

		binaryFind(arr, leftIdx, midIdx-1, val)

	} else if (*arr)[midIdx] < val {

		binaryFind(arr, midIdx+1, rightIdx, val)

	} else {
		fmt.Printf("找到了在%v位\n", midIdx)
	}

}

//查找
//1. 順序查找
//2. 二分查找(前提是陣列是要排序好的)  重要!!!!

//二分查找 (一直折半，很像終極密碼)
//有2個假想指針  左指針、右指針
//   左指針 : 指向陣列首位 (leftIdx)
//   右指針 : 指向陣列末位 (rightIdx)
//1. 先找到陣列中間的index =>  (leftIdx + rightIdx)/2    (midIdx)
//2. 定義要找的值 val ，與arr[midIdx] 進行比較
//if  arr[midIdx] > val ，代表他在arr[midIdx]左邊 => 查找範圍 leftIdx ~~~ arr[midIdx]-1
//if  arr[midIdx] < val ，代表他在arr[midIdx]右邊 => 查找範圍 arr[midIdx]+1 ~~~ rightIdx
//if  arr[midIdx] = val ，找到了

//3. 退出遞迴條件 =>左右標會不斷靠近，當 左>右 還找不到就代表 真的找不到了
//if leftIdx > rightIdx ，代表找不到
