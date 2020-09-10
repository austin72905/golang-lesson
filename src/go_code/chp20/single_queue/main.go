package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	//new 一個queue
	//struct初始化時，陣列也初始化了
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	//這邊取出來元素之後，前面的空格不能用了
	var key string
	var val int
	for {
		fmt.Println("1. 輸入add 添加元素到陣列")
		fmt.Println("2. 輸入get 從陣列獲取數據")
		fmt.Println("3. 輸入show 顯示陣列")
		fmt.Println("4. 輸入exit 退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("輸入數字加入列隊")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入成功")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("從列隊取出一個數", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}

//Queue struct
type Queue struct {
	maxSize int
	array   [5]int //陣列=>模擬queue
	front   int    //首指針
	rear    int    //尾指針
}

//AddQueue is metod belongs to Queue to add val to Queue Arr
//將數據存入列隊 addQueue
func (que *Queue) AddQueue(val int) (err error) {
	//先判斷queue是否已滿
	if que.rear == que.maxSize-1 {
		//rear 是列隊尾部(含最後元素)
		return errors.New("queue full")
	}

	que.rear++ //rear 後移
	que.array[que.rear] = val
	return
}

//ShowQueue is metod belongs to Queue to show elements in array
//showQueue 顯示列隊
func (que *Queue) ShowQueue() {

	fmt.Println("列隊當前情況")
	//顯示列隊，找到隊首，遍歷到隊尾
	//front 在隊首，不包含第一個元素
	for i := que.front + 1; i <= que.rear; i++ {
		fmt.Printf("array[%d]=%d ", i, que.array[i])
	}
	fmt.Println()
}

//GetQueue is metod belongs to Queue to get element from array
//getQueue  從列隊取出數據
func (que *Queue) GetQueue() (val int, err error) {
	//判斷列隊是否空
	if que.front == que.rear {
		return -1, errors.New("queue empty") //在main判斷的時候-1這個值其實也用不用
	}
	//要取之前，先把頭指針往後移
	que.front++
	//只到要取出的元素的index，在取出來，這樣還是維持front再隊首，但不包含第一個元素
	val = que.array[que.front]
	return val, err
}

// 列隊 queue
// 1. 先進先出
// 2. 可用鏈結串列、陣列表示  P.S.陣列取出元素後，前面的空間會空出來，如何應用空間?
//     (1)使用陣列
// 	a.要聲明MaxSize
// 	b.需要兩個變數 front、rear 標記列隊頭尾的index => front 指向最前面，但不包含當前最頭的元素 (front 、 rear 起始值 -1)

//實作一個非環形的列隊
//將數據存入列隊 addQueue
//getQueue  從列隊取出數據
//showQueue 顯示列隊
//1. 將尾指針 後移: rear+1 , 當front==rear ,代表列隊空了
//2. rear== MaxSize-1(最大index)，代表列隊滿了

//這種非環狀的列隊，元素取出來的空格，無法重複利用，只是一次性的queue
