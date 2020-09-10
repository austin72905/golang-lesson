package main

import "fmt"

func main() {
	//1. 定義一個陣列
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑旗
	chessMap[2][3] = 2 //白棋

	//2. 輸出
	//index =一維陣列
	for _, val := range chessMap {
		for _, val2 := range val {
			fmt.Printf("%d ", val2) //可用/t 隔開
		}
		fmt.Println()
	}

	//印出來很多0，浪費空間
	//3. 轉成稀疏陣列
	//(1)遍歷chessMap，如果發現元素不為0，創建node結構體
	//(2)將其放入對應的切片中

	var sparseArr []ValNode
	//golang 聲明陣列時，大小寫死，而標準的稀疏陣列，要記錄二維陣列(行和列，默認值)
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}

	sparseArr = append(sparseArr, valNode)

	for inx0, val0 := range chessMap {
		for inx1, val1 := range val0 {
			//創建新節點
			if val1 != 0 {
				valNode := ValNode{
					row: inx0,
					col: inx1,
					val: val1,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	fmt.Println("稀疏數組")
	//輸出
	for index, valNode := range sparseArr {
		fmt.Printf("%d => %d %d %d\n", index, valNode.row, valNode.col, valNode.val)
	}

	//將稀疏陣列，轉回原來的陣列
	//1. 定義一個陣列
	var revChessMap [11][11]int

	//2. 例遍 稀疏陣列
	for index, valNode := range sparseArr {
		//第0 row 是 11 11 0 =>只是記錄原始陣列大小的
		if index != 0 {
			revChessMap[valNode.row][valNode.col] = valNode.val
		}

	}

	fmt.Println("還原後的原始陣列")
	//輸出
	for _, val := range revChessMap {
		for _, val2 := range val {
			fmt.Printf("%d ", val2) //可用/t 隔開
		}
		fmt.Println()
	}

}

//ValNode struct
type ValNode struct {
	row int
	col int
	val int
}
