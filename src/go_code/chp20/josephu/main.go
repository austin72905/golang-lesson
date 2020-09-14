package main

import "fmt"

func main() {
	first := addNode(5)
	//顯示
	showNdoe(first)

	//開始報數
	playGame(first, 2, 3)
}

//Boy struct
type Boy struct {
	no   int  //編號
	next *Boy //指向下一個節點

}

//寫一個函數構成單向環形練表
//num : 節點數
//*Boy : 返回鍊表第一個小孩的值指針
func addNode(num int) *Boy {

	first := &Boy{}
	////輔助指針
	current := &Boy{}

	if num < 1 {
		fmt.Println("num 不得小於1")
		return first
	}

	//循環構建環形練表
	for i := 1; i <= num; i++ {

		boy := &Boy{
			no: i,
			//不給他next 值，默認nil
		}
		//輔助指針

		//1. 第一個小孩比較特殊
		if i == 1 {
			//讓頭指針指向他
			first = boy //頭指針不要動
			//輔助指針
			current = boy
			//只有一個指針也要形成環形
			current.next = first
		} else {
			current.next = boy
			//current 往後移一個
			current = boy
			//構成環形練表
			current.next = first
		}
	}
	return first

}

//開始報數
func playGame(first *Boy, startN0 int, couuntNum int) {

	//如果是空鍊表
	if first.next == nil {
		fmt.Println("空的鍊表")
		return
	}

	//開始的編號>小孩總數

	//輔助指針
	tail := first
	//讓tail指向最後一個 => 這樣 tail 往前時就會比temp 慢一步 tail 下一個是head ，temp又指向first
	for {

		//已經到最後了
		if tail.next == first {
			break
		}

		tail = tail.next
	}

	//讓first移動到startN0 (後面刪除小孩已first 為主)
	//no 從1 開始
	for i := 1; i <= startN0-1; i++ {
		first = first.next
		tail = tail.next
	}

	fmt.Println()
	//數couuntNum，開始刪除first 指向的小孩
	for {
		//開始數
		//自己也要報數所以移動 couuntNum-1
		for i := 1; i <= couuntNum-1; i++ {
			first = first.next
			tail = tail.next
		}

		//刪除first指向的節點
		//1. 先讓first往前走一步
		//2. 讓tail.next指向first (跳過中間那個，沒人指向他就被刪除了)
		fmt.Printf("男孩出列順序為 no %d\n", first.no)
		first = first.next
		tail.next = first

		//退出條件 : 只剩一個小孩 tail ==first
		if tail == first {
			break
		}

	}

	fmt.Printf("最後出列的小孩no %d\n", first.no)
}

//顯示環形練表
func showNdoe(first *Boy) {

	//如果環形練表回空
	if first.next == nil {
		fmt.Println("空鍊表")
		return
	}

	//創建一個輔助指針
	current := first
	for {
		fmt.Printf("男孩編號=%d ->", current.no)
		//退出條件
		//已經到最後一個了
		if current.next == first {
			break
		}
		//移動到下一個
		current = current.next
	}
}

//約瑟夫問題

//編號 1~n 的人圍成一圈，並指定從編號 k 的人開始報數(1 <= k <= n)，數到m的人出列
//下一位繼續從1開始報數，直到所有人出列

//思路 : 環形鍊表

//實際應用
//1. 5個人圍成一圈  n=5

//2. 從編號2的人開始報數 k=2

//3. 報數道第3個人時，該人出列  m=3
