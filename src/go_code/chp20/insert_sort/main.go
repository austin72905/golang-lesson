package main

import "fmt"

func main() {

	arr := [5]int{23, 0, 12, 56, 34}

	//分解動作
	//insertSort(&arr)

	//用for循環
	insertSortP(&arr)
}

//大到小
func insertSort(arr *[5]int) {

	//第一個元素已經放入有序組，所以從第二個開始比較

	//第一次
	//無序組要比較的元素
	//把它存在變數裡面，後面數組的植被覆蓋掉也沒關係
	insertVal := arr[1]
	//比較的位置 (都是在插入的前一個)
	cmptIndex := 1 - 1

	for cmptIndex >= 0 && arr[cmptIndex] < insertVal {
		arr[cmptIndex+1] = arr[cmptIndex] //後移 (將比較的元素值"複製"，在比較的位置(cmptIndex)的下一個(cmptIndex + 1)插入)

		cmptIndex-- //一直往前移，小於 0 就會退出了
	}

	//插入
	//代表cmptIndex 有往前移過   前面 cmptIndex = 1 - 1 ，如果cmptIndex 都沒有移過 ，cmptIndex + 1 = 1
	if cmptIndex+1 != 1 {
		arr[cmptIndex+1] = insertVal
	}

	fmt.Println("第一次插入 ", *arr)

	//第二次
	//無序組要比較的元素
	insertVal = arr[2]
	//比較的位置 (都是在插入的前一個)
	cmptIndex = 2 - 1

	for cmptIndex >= 0 && arr[cmptIndex] < insertVal {
		arr[cmptIndex+1] = arr[cmptIndex] //後移 (將比較的元素值"複製"，在比較的位置(cmptIndex)的下一個(cmptIndex + 1)插入)

		cmptIndex-- //一直往前移，小於 0 就會退出了
	}

	//插入
	//代表cmptIndex 有往前移過   前面 cmptIndex = 2 - 1 ，如果cmptIndex 都沒有移過 ，cmptIndex + 1 = 2
	//
	if cmptIndex+1 != 2 {
		arr[cmptIndex+1] = insertVal
	}

	fmt.Println("第二次插入 ", *arr)

	//第三次
	//無序組要比較的元素
	insertVal = arr[3]
	//比較的位置 (都是在插入的前一個)
	cmptIndex = 3 - 1

	for cmptIndex >= 0 && arr[cmptIndex] < insertVal {
		arr[cmptIndex+1] = arr[cmptIndex] //後移 (將比較的元素值"複製"，在比較的位置(cmptIndex)的下一個(cmptIndex + 1)插入)

		cmptIndex-- //一直往前移，小於 0 就會退出了
	}

	//插入
	//代表cmptIndex 有往前移過   前面 cmptIndex = 3 - 1 ，如果cmptIndex 都沒有移過 ，cmptIndex + 1 = 3
	//
	if cmptIndex+1 != 3 {
		arr[cmptIndex+1] = insertVal
	}

	fmt.Println("第三次插入 ", *arr)

	//第四次
	//無序組要比較的元素
	insertVal = arr[4]
	//比較的位置 (都是在插入的前一個)
	cmptIndex = 4 - 1

	for cmptIndex >= 0 && arr[cmptIndex] < insertVal {
		arr[cmptIndex+1] = arr[cmptIndex] //後移 (將比較的元素值"複製"，在比較的位置(cmptIndex)的下一個(cmptIndex + 1)插入)

		cmptIndex-- //一直往前移，小於 0 就會退出了
	}

	//插入
	//代表cmptIndex 有往前移過   前面 cmptIndex = 4 - 1 ，如果cmptIndex 都沒有移過 ，cmptIndex + 1 = 4
	//
	if cmptIndex+1 != 4 {
		arr[cmptIndex+1] = insertVal
	}

	fmt.Println("第四次插入 ", *arr)

}

//大到小
func insertSortP(arr *[5]int) {

	//第一個元素已經放入有序組，所以從第二個開始比較

	//最後一個員數要參與比較，所以不用-1
	for round := 1; round < len(arr); round++ {
		//無序組要比較的元素
		//把它存在變數裡面，後面數組的植被覆蓋掉也沒關係
		insertVal := arr[round]
		//比較的位置 (都是在插入的前一個)
		cmptIndex := round - 1

		for cmptIndex >= 0 && arr[cmptIndex] < insertVal {
			arr[cmptIndex+1] = arr[cmptIndex] //後移 (將比較的元素值"複製"，在比較的位置(cmptIndex)的下一個(cmptIndex + 1)插入)

			cmptIndex-- //一直往前移，小於 0 就會退出了
		}

		//插入
		//代表cmptIndex 有往前移過   前面 cmptIndex = 1 - 1 ，如果cmptIndex 都沒有移過 ，cmptIndex + 1 = 1，就不用代替值了，代表他在有序組中是合適的位置
		if cmptIndex+1 != round {
			arr[cmptIndex+1] = insertVal
		}

		fmt.Printf("第%d次插入 %v\n", round, *arr)
	}

}

//插入排序法
//將要排序的 有 n 個數的數組 分成兩堆，一個有序，一個是待排序

//需要有一個比較的指針 cmptIndex ，插入的位置在 insertIndex + 1

//將無序組 中要比較的元素 存在一個 insertVal

//將第一個元素 放入 有序 組 => 有序組: 1 個  待排序組 : n-1 個

//將待排序組第一個元素取出， 逐一跟有序組的元素比較，如果找到適合的位置就插入

//如果沒有，就將比較的元素值"複製"，在比較的位置(cmptIndex)的下一個(cmptIndex + 1)插入 (形成移動的效果)

//ex: 要比較的數組 [23,0,12,56,34] 大到小排序

//1. [23]       [0,12,56,34]

//2. [23,0]       [12,56,34]

//3. [23,0,0]        [56,34]
//   [23,12,0]       [56,34]

//4. [23,12,0,0]        [34]
//   [23,12,12,0]       [34]
//   [23,23,12,0]       [34]
//   [56,23,12,0]       [34]

//5. [56,23,12,0,0]
//   [56,23,12,12,0]
//   [56,23,23,12,0]
//   [56,34,23,12,0]

// 比較指針 insetIndex 會從 有序組中，最後一個逐一往 前面移動
