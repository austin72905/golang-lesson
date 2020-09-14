package main

import "fmt"

//HeroNode struct
//定義一個英雄節點
type HeroNode struct {
	no       int
	name     string
	nickName string
	next     *HeroNode //指向下一個節點
}

//給鏈結串列插入一個節點
//第一種方式:在鏈結串列最後加入
//要先找到頭節點
func addNode(head *HeroNode, newNode *HeroNode) {
	//1. 先找到最後的節點
	//2. 創建一個輔助節點
	temp := head
	for {
		//代表已經到最後了
		if temp.next == nil {
			break
		}
		//讓temp不斷指向下一個節點
		temp = temp.next
	}
	//3. 將newNode 加入鏈結串列的最後
	temp.next = newNode
}

//第二種方式:根據no 的編號，從小到大
//要先找到頭節點
func addNodeByNo(head *HeroNode, newNode *HeroNode) {
	//1. 要找到欲插入節點的下一個(temp、temp.next，欲插入的節點插在這兩個中間)
	//2. 創建一個輔助節點
	temp := head
	flag := true //標示
	for {
		//判斷是否到鏈結的最後
		if temp.next == nil {
			break
		} else if temp.next.no > newNode.no {
			//代表找到插入點了
			//如果要no大排到小，改成<
			//看業務邏輯，如果no相等可加>=
			break
		} else if temp.next.no == newNode.no {
			//有相同no就不給加入了
			flag = false
			break
		}
		//讓temp指到下一個
		temp = temp.next
	}

	if !flag {
		fmt.Println("have same no", newNode.no)
		return
	}
	//插入節點
	//先讓新節點的下一個指向temp.next
	newNode.next = temp.next
	temp.next = newNode

}

//刪除節點
func deleteNode(head *HeroNode, id int) {
	temp := head
	flag := false //標示
	//找到要刪除的節點
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			//代表找到要刪出的節點了
			flag = true
			break
		}
		//讓temp指向下一個節點
		temp = temp.next
	}

	if flag {
		//讓temp的下一個節點，指向下下個節點(因為下個節點是要被刪掉了，只要沒人指向他，他就會被GC)
		temp.next = temp.next.next
	} else {
		fmt.Println("id is not exist")
	}
}

//顯示鏈結串列所有訊息
func showNode(head *HeroNode) {
	//1. 創建一個輔助節點
	temp := head

	//先判斷是不是空鏈結
	if temp.next == nil {
		fmt.Println("it is empty linklist")
		return
	}
	//2.遍歷這個節點
	for {
		//因為已經確認下一個節點不會空了
		fmt.Printf("[%d , %s ,%s]->", temp.next.no, temp.next.name, temp.next.nickName)

		temp = temp.next
		//判斷是否到最尾了
		if temp.next == nil {
			break
		}
	}
}

func main() {
	//1. new 一個頭節點
	//頭節點不需要給值
	head := &HeroNode{}

	//2. 創建一個新的herrNdoe
	hero1 := &HeroNode{
		no:       1,
		name:     "琦玉",
		nickName: "一拳超人",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "傑諾斯",
		nickName: "超級生化人",
	}
	//3. 加入
	//這個寫法，只能將新節點加到最後，如果要插入中間就GG了
	// addNode(head, hero1)
	// addNode(head, hero2)

	//按照no 排序
	addNodeByNo(head, hero2)
	addNodeByNo(head, hero1)
	addNodeByNo(head, hero1)
	//4. 顯示
	showNode(head)

	fmt.Println()
	//5. 刪除
	deleteNode(head, 1)
	showNode(head)
}

//1. 第一種方式:在鏈結串列最後加入
//2. 第二種方式:根據no 的編號，從小到大

//單向鏈結串列
//1. 查找只能是一個方向
//2. 需要依靠輔助節點
//雙向鏈結串列
//1. 查找可以雙向
//2. 可以自我刪除
