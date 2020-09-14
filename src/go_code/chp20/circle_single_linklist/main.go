package main

import "fmt"

//CatNode struct
type CatNode struct {
	no   int
	name string
	next *CatNode
}

func main() {
	//初始化一個頭節點
	head := &CatNode{}

	//初始化一個新節點
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}

	//
	cat2 := &CatNode{
		no:   2,
		name: "tina",
	}

	cat3 := &CatNode{
		no:   1,
		name: "bob",
	}

	addCatNode(head, cat1)
	addCatNode(head, cat2)
	addCatNode(head, cat3)
	showNode(head)
	head = deleteNode(head, 1)
	showNode(head)
}

func addCatNode(head *CatNode, newNode *CatNode) {

	//判斷是不是添加第一個節點
	if head.next == nil {
		head.no = newNode.no
		head.name = newNode.name
		//head.next = newNode //形成一個環形
		head.next = head //要指向自己才是環狀
		//只有一個節點也是環形(自己指向自己)
		fmt.Println(newNode, "加入環形")
		return
	}

	//定義一個臨時變數，協助找到環形的最後節點
	temp := head

	for {
		if temp.next == head {
			//代表找到最後一個了
			break
		}
		temp = temp.next //把temp往後移
	}

	//加入鏈結串列
	temp.next = newNode
	//因為是環狀，所以要讓加入的節點指向頭節點
	newNode.next = head

}

//輸出環形練表
func showNode(head *CatNode) {
	temp := head

	if temp.next == nil {
		//如果裡面有一個節點，他會指向(自己)
		fmt.Println("目前是空的練表")
		return
	}

	for {
		fmt.Printf("喵的訊息=[no=%d  name=%s] ->\n", temp.no, temp.name) //如果把temp整個打印出來，因為裡面有next指向自己，會報錯
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

//刪除
//現在頭節點也有值了，也要進行比較
//如果刪掉的剛好是頭節點，原本傳進去的head就不見了，所以用一個回傳值
//這個方法只能刪第一個找到的no 不能重複，刪掉一個他就跳出了
func deleteNode(head *CatNode, id int) *CatNode {

	temp := head
	helper := head
	if temp.next == nil {
		fmt.Println("空的列表")
		return head
	}

	//將helper定位到最後 => 這樣 helper 往前時就會比temp 慢一步 helper 下一個是head ，temp又指向head
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	//如果只有一個節點
	if temp.next == head {
		temp.next = nil
		return head
	}
	flag := true
	//如果兩個以上節點
	for {
		//已經到最後還沒找到(但最後一個還沒比較)
		if temp.next == head {
			break
		}
		//考慮是不是刪到頭節點
		if temp.no == id {
			//剛好刪到頭節點
			if temp.no == id {
				head = head.next
			}
			//找到要刪除的了
			helper.next = temp.next //找到了temp，把它刪掉
			fmt.Printf("節點no=%d\n", id)
			flag = false
			break
		}

		temp = temp.next     // temp 是用來比較 是不是 = 要刪除的id
		helper = helper.next // helper 是用來刪的
	}
	//這邊還要比較一次
	//如果temp為真，代表都還沒刪除過
	if flag {
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("節點no=%d\n", id)
		} else {
			fmt.Println("節點不存在")
		}
	}

	return head
}

//雖然環形沒有頭節點，但還是要自己初始化一個，不然沒辦法開始

//刪除環形練表思路

//1.先讓temp 指向 head

//2.讓helper 指向環形練表的最後

//3.讓temp 和要刪除的id 進行比較，如果相同，則同helper 完成刪除(要考慮如果刪掉的是頭節點怎麼辦)
