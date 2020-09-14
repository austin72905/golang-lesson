package main

import "fmt"

func main() {
	//定義一個陣列
	arr := [5]int{10, 34, 19, 100, 80}

	//一次一次排
	//selectSort(&arr)

	//使用for循環
	selectSortP(&arr)
}

//完成排序
//大到小
func selectSort(arr *[5]int) {
	//標準訪問方式
	//(*arr)[1] = 600

	//這樣也可以改變
	//arrp[1]=600

	//第一次交換
	//1. 先將最大值，與arr[0]交換
	//假設arr[0] 是最大值
	max := arr[0]
	maxInx := 0

	//2. 遍歷
	//假設arr[0]，最大了，所以從下一個開始比較
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			//找到最大值
			//不是在這邊交換，最後才交換效率才會高
			max = arr[i]
			maxInx = i
		}
	}

	//3. 交換
	if maxInx != 0 {
		arr[0], arr[maxInx] = arr[maxInx], arr[0]
	}

	fmt.Println("第一次", *arr)

	//第二次交換
	//1. 先將最大值，與arr[1]交換
	//假設arr[1] 是最大值
	max = arr[1]
	maxInx = 1

	//2. 遍歷
	//假設arr[1]，最大了，所以從下一個開始比較
	for i := 2; i < len(arr); i++ {
		if max < arr[i] {
			//找到最大值
			//不是在這邊交換，最後才交換效率才會高
			max = arr[i]
			maxInx = i
		}
	}

	//3. 交換
	if maxInx != 1 {
		arr[1], arr[maxInx] = arr[maxInx], arr[1]
	}

	fmt.Println("第二次", *arr)

	//第三次交換
	//1. 先將最大值，與arr[2]交換
	//假設arr[2] 是最大值
	max = arr[2]
	maxInx = 2

	//2. 遍歷
	//假設arr[2]，最大了，所以從下一個開始比較
	for i := 3; i < len(arr); i++ {
		if max < arr[i] {
			//找到最大值
			//不是在這邊交換，最後才交換效率才會高
			max = arr[i]
			maxInx = i
		}
	}

	//3. 交換
	if maxInx != 2 {
		arr[2], arr[maxInx] = arr[maxInx], arr[2]
	}

	fmt.Println("第三次", *arr)

	//第四次交換
	//1. 先將最大值，與arr[3]交換
	//假設arr[3] 是最大值
	max = arr[3]
	maxInx = 3

	//2. 遍歷
	//假設arr[3]，最大了，所以從下一個開始比較
	for i := 4; i < len(arr); i++ {
		if max < arr[i] {
			//找到最大值
			//不是在這邊交換，最後才交換效率才會高
			max = arr[i]
			maxInx = i
		}
	}

	//3. 交換
	if maxInx != 3 {
		arr[3], arr[maxInx] = arr[maxInx], arr[3]
	}

	fmt.Println("第四次", *arr)
}

//完成排序
//從 selectSort 可以發現代碼很多重複的
func selectSortP(arr *[5]int) {
	//標準訪問方式
	//(*arr)[1] = 600

	//這樣也可以改變
	//arrp[1]=600

	//如果有5個元素要排序4次... 推出 如果有n 個元素，最多要排序n-1 次

	for round := 0; round < len(arr)-1; round++ {
		//第一次交換
		//1. 先將最大值，與arr[0]交換
		//假設arr[0] 是最大值
		max := arr[round]
		maxInx := round

		//2. 遍歷
		//假設arr[0]，最大了，所以從下一個開始比較
		for i := round + 1; i < len(arr); i++ {
			if max < arr[i] {
				//找到最大值
				//不是在這邊交換，最後才交換效率才會高
				max = arr[i]
				maxInx = i
			}
		}

		//3. 交換
		if maxInx != round {
			arr[round], arr[maxInx] = arr[maxInx], arr[round]
		}

		fmt.Printf("第%d次 %v\n", round+1, *arr)
	}

}

//冒泡排序:
//依次和相鄰的元素比較，使較小的元素往前交換(陣列單元有講過)

//選擇排序:
//有1~n個元素
//第一次:從1~n ，找出裡面 最小(或最大)的元素，與第1個元素交換
//第二次:從2~n...找出裡面 最小(或最大)的元素，與第2個元素交換
//第三次:從3~n...找出裡面 最小(或最大)的元素，與第3個元素交換
