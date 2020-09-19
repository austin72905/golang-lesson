package main

import (
	"errors"
	"fmt"
)

func main() {

	//使用陣列，模擬stack
	stack := &stack{
		maxTop: 5,  //最多存放5個數
		top:    -1, //-1 表示空棧
	}

	//加入元素到stack
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(5)

	//顯示
	stack.show()

	//出棧
	val, _ := stack.pop()
	fmt.Println("取出的元素 =", val)

	stack.show()
}

type stack struct {
	maxTop int // 棧頂最大值
	top    int // 棧頂
	//可以不用寫棧底，因為其是固定的 0 or -1

	arr [5]int //陣列
}

//新增元素到stack
func (sta *stack) push(val int) (err error) {

	//先判斷是否滿了
	if sta.top == sta.maxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}

	sta.top++ //因為初始化是-1

	sta.arr[sta.top] = val
	return

}

//遍歷棧 (要從棧頂開始)
func (sta *stack) show() {

	//先判斷stack是否為空
	if sta.top == -1 {
		fmt.Println("stack empty")
		return
	}

	//遍歷(從棧頂)
	for i := sta.top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, sta.arr[i])
	}
}

//出棧
//回傳兩個值 1. 取出的元素 2. 錯誤
func (sta *stack) pop() (val int, err error) {
	//判斷是否為空
	if sta.top == -1 {
		fmt.Println("stack empty")
		return -1, errors.New("stack empty")
	}

	//先取值
	val = sta.arr[sta.top]
	sta.top--
	return val, nil
}

//棧(stack)

//特性
//1. 先進後出 (跟彈夾一樣)

//2. 刪除和插入元素，都在同一端 棧頂(top) ，固定的一端 棧底(bottom)

//3. 會有兩個指針 top 、 bottom ，如果是空stack  ， top、bottom都 = -1

//4. 新增了一個元素 top = 0 , bottom = 0

//5. 再新增一個元素 top = 1 , bottom = 0

//6. 當 top = len(array)-1 ，代表stack 滿了

//應用場景

//1. 子程序的調用 ，會將下個指令的地址存入stack，直到子程序執行完，再把地址取出，以回到原本的程序
//2. 處理遞迴
//3. 表達式的轉換
//4. 二叉樹遍歷
//5. 圖形深度優先搜索法
