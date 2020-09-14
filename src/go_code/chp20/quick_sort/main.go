package main

import "fmt"

func main() {
	arr := [6]int{-9, 78, 0, 23, -567, 70}

	//第一次排序
	//quickSort(0, len(arr)-1, &arr)

	//完整的快速排序法
	quickSortP(0, len(arr)-1, &arr)
	fmt.Println(arr)
}

//快速排序
//left 數組最左邊的index
//right 數組最右邊的index
func quickSort(left int, right int, arr *[6]int) {

	//左指針，會不停向右移動 ，直到找到比基準點大的數
	lptr := left
	//右指針，會不停向左移動 ，直到找到比基準點小的數
	rptr := right

	//基準點 ，取中間的數
	point := arr[(left+right)/2]

	//把比 point  小的都放左邊
	//比   point  大的放右邊
	//lptr、rptr會不斷接近，lptr < rptr代表整個數組都比較過了
	//比point 小的 都在左邊 ;  比point 大的  都在右邊
	//[-9 -567 0 23 78 70]
	for lptr < rptr {

		//從point 左邊找到 >= point 的 值
		for arr[lptr] < point {
			lptr++
		}
		//從point 右邊找到<= point 的 值
		for arr[rptr] > point {
			rptr--
		}

		//{-9, 78, 0 ,23, -567, 70}  point=arr[(0+5)/2]=> 0
		// lptr 找到的數78  ;  rptr 找到的數 -567
		//找到以後把他們兩個交換

		//代表兩個指針找完了
		if lptr >= rptr {
			break
		}

		//如果找到了，把lptr 跟  rptr 指到的值交換
		//{-9, -567, 0 ,23, 78, 70}
		arr[lptr], arr[rptr] = arr[rptr], arr[lptr]

		//小優化
		//如果剛好數組是基數，lptr 跟  rptr 剛好都只到最中間的時候，跳過不用比較
		if arr[lptr] == point {
			rptr--
		}
		if arr[rptr] == point {
			lptr++
		}

	}
}

//快速排序(有遞迴)
//left 數組最左邊的index
//right 數組最右邊的index
func quickSortP(ftIdx int, ltIdx int, arr *[6]int) {

	//左指針，會不停向右移動 ，直到找到比基準點大的數
	lptr := ftIdx
	//右指針，會不停向左移動 ，直到找到比基準點小的數
	rptr := ltIdx

	//基準點 ，取中間的數
	point := arr[(ftIdx+ltIdx)/2]

	//把比 point  小的都放左邊
	//比   point  大的放右邊
	//lptr、rptr會不斷接近，lptr < rptr代表整個數組都比較過了
	//比point 小的 都在左邊 ;  比point 大的  都在右邊
	//[-9 -567 0 23 78 70]
	for lptr < rptr {

		//如果要排大到小， 把 <、>互換即可
		//從point 左邊找到 >= point 的 值
		for arr[lptr] < point {
			lptr++
		}
		//從point 右邊找到<= point 的 值
		for arr[rptr] > point {
			rptr--
		}

		//{-9, 78, 0 ,23, -567, 70}  point=arr[(0+5)/2]=> 0
		// lptr 找到的數78  ;  rptr 找到的數 -567
		//找到以後把他們兩個交換

		//代表兩個指針找完了
		if lptr >= rptr {
			break
		}

		//如果找到了，把lptr 跟  rptr 指到的值交換
		//{-9, -567, 0 ,23, 78, 70}
		arr[lptr], arr[rptr] = arr[rptr], arr[lptr]

		//小優化
		//如果剛好數組是基數，lptr 跟  rptr 剛好都只到最中間的時候，跳過不用比較
		if arr[lptr] == point {
			rptr--
		}
		if arr[rptr] == point {
			lptr++
		}

	}

	//剛剛跳出for迴圈 ，目前可以確定 lptr >= rptr
	//不做這步會卡死
	//因為這個兩個就會剛好都指向point，走下面的遞迴，會無限月讀
	if lptr == rptr {
		lptr++
		rptr--
	}

	//目前可以確定 lptr > rptr

	//遞迴
	//向左遞迴 (對 point 左半邊的數組做遞迴)
	if ftIdx < rptr {
		quickSortP(ftIdx, rptr, arr)
	}
	//向右遞迴 (對 point 右半邊的數組做遞迴)
	if ltIdx > lptr {
		quickSortP(lptr, ltIdx, arr)
	}
}

//快速排序法

//小到大排序
//從一堆數組中，找一個數當基準點， 比該數小=>放到該數左邊 ; 比該數大=>放到該數右邊
//用遞迴 ，對左邊、跟右邊的數列，再做上次的動作，直到整個數列排序正確

//ex:  2 , 10 , 8 , 22 , 34 , 5 , 12 , 28 , 21 , 11   一般來說都取中間數為基準，但不一定  ... 11 為基準

//     2 , 10 , 8 , 5   [11]   22 , 34 , 12 , 28 , 21     左邊...5 為基準  ;  右邊...21 為基準

//     2  [5]  10 , 8          12  [21]  22 , 28 , 34     左邊...8 為基準  ;  右邊...34為基準

//            [8]  10                    22 , 28  [34]   左邊..都排序好了 ;  右邊...28為基準

//                                       22  [28]

//最後排序結果    2,5,8,10,11,12,21,22,28,34
