package main

import (
	"errors"
	"fmt"
	"os"
)

//CircleQuere struct
type CircleQuere struct {
	maxSize int
	array   [5]int
	head    int //頭指針
	tail    int //尾指針
}

func main() {

	//new一個環形隊列
	queue := &CircleQuere{
		maxSize: 5, //實際上也只能裝4個
		head:    0,
		tail:    0,
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

//AddQueue is method belongs to Queue to add val to Queue Arr
//將數據存入列隊 addQueue
func (que *CircleQuere) AddQueue(val int) (err error) {
	//判斷容器是否滿了
	if que.IsFull() {
		return errors.New("queue full")
	}
	//先放值加入尾部，tail 不包含最後一個元素，通常都指到最後一個元素的下一個
	que.array[que.tail] = val //把值留給尾部
	//que.tail++ 這樣寫會報錯，因為是環狀的隊列
	que.tail = (que.tail + 1) % que.maxSize
	return
}

//ShowQueue is method belongs to Queue to show elements in array
//showQueue 顯示列隊
func (que *CircleQuere) ShowQueue() {

	fmt.Println("環形隊列情況如下")

	//不能單純地用遍歷了，因為是環狀，有可能head 在tail後面
	size := que.Size()
	if size == 0 {
		fmt.Println("queue is empty")
	}

	//設計一個輔助變數，指向head 所在index
	tempHead := que.head

	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d ", tempHead, que.array[tempHead])
		//因為是環形的，\才可以又跑到頭
		tempHead = (tempHead + 1) % que.maxSize //不是size喔
	}
	fmt.Println()
}

//GetQueue is method belongs to Queue to get element from array
//getQueue  從列隊取出數據
func (que *CircleQuere) GetQueue() (val int, err error) {
	if que.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	//取出
	//head 指向隊首，且包含第一個元素
	val = que.array[que.head]
	//que.head++ 這樣寫會報錯
	que.head = (que.head + 1) % que.maxSize
	return
}

//IsFull is method belongs to Queue to check arr is full or not
//判斷是否容器是否滿了
func (que *CircleQuere) IsFull() bool {
	return (que.tail+1)%que.maxSize == que.head
}

//IsEmpty is method belongs to Queue to check arr is empty or not
//判斷是否容器是否為空
func (que *CircleQuere) IsEmpty() bool {
	return que.tail == que.head
}

//Size is method belongs to Queue to get count of elements in arr
//取出列隊有多少元素
func (que *CircleQuere) Size() int {
	//也是有可能head 在 tail 後面
	return (que.tail + que.maxSize - que.head) % que.maxSize
}

//環形列隊

//1. 將容器空出一格，作為約定來做以下判斷

//1-1. 因為是環狀，當尾指針到尾了，要再回到最前面，所以有以下公式   (當tail指到index 4 ，代表實際上這個容器是[5])，所以要+1

//2. (tail + 1) % maxSize ==head 代表隊列滿了

//3. tail == head 代表隊列為空

//4. 初始化時 tail = 0 、 head = 0

//5. 如何統計有幾個元素?  ( tail + maxSize - head ) %  maxSize    => 因為它會空出一格(隨時保持一隔空出，不一定是第幾格)，當tail 指到 index 1 時，實際上只有一個元素

//6. tail 不包含最後一個元素，通常都指到最後一個元素的下一個

//7. head 指向隊首，且包含第一個元素
