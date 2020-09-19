package main

import (
	"errors"
	"fmt"
	"strconv"
)

//對於ch 再做一些優化
//處理多位數 =>主要是處理掃描到數的時候，要再往index後探一位
//1. 如果到尾端，就不往後了
//2. 如果index後一位是 "運算符" 把數加入numStack吧
//3. 把數拼接起來 ，再加入numStack
func main() {

	//創建兩個棧(數棧)numStack、(運算符棧)operStack
	numStack := &stack{
		maxTop: 20,
		top:    -1,
	}

	operStack := &stack{
		maxTop: 20,
		top:    -1,
	}

	//表達式
	exp := "30+2*6-2"
	//定義一個index ，遍歷掃苗exp
	index := 0

	//運算的變數
	num1 := 0
	num2 := 0
	oper := 0
	result := 0

	//定義一個變數拼接數字
	keepNum := ""

	for {

		//處理多位數 =>主要是處理掃描到數的時候，要再往index後探一位
		//1. 如果到尾端，就不往後了
		//2. 如果index後一位是 "運算符" 把數加入numStack吧
		//3. 把數拼接起來 ，再加入numStack

		ch := exp[index : index+1] //掃描exp ，取出來的是單一個字的"字串"
		//轉成acsii碼
		temp := int([]byte(ch)[0]) //先轉成byte，但是方法都是要傳int進去，所以再轉成int

		if operStack.isOper(temp) {
			//如果是運算符
			//如果operStack 是空stack
			if operStack.top == -1 {
				operStack.push(temp)
			} else {
				//如果不是 空stack
				//     a. 如果operStack "頂端的" 符號  ， 優先運算  於  "準備加入的" 符號，  (先*/後+-)
				//        從 operStack 取出 符號 ， numStack 也取出兩個數 來運算，
				//        運算完把結果重新放到 numStack  ，"準備加入的" 運算符 ，再加入operStack
				if operStack.prior(operStack.arr[operStack.top]) >= operStack.prior(temp) {

					num1, _ = numStack.pop()
					num2, _ = numStack.pop()
					oper, _ = operStack.pop()
					result = operStack.cal(num1, num2, oper)

					numStack.push(result)
					operStack.push(temp)
				} else {
					//b. 如果operStack "準備加入的" 符號 ， 優先運算  於  "頂端的" 符號 => 運算符 直接加入operStack
					operStack.push(temp)
				}
			}
		} else {
			//如果是數字 => numStack
			//要先把acsii 碼 轉成 原本表達式的數

			//拼接
			keepNum += ch

			//處理多位數
			//如果到尾端，就不往後探了
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.push(int(val))
			} else {
				//如果index後一位是 "運算符" 把數加入numStack吧
				if operStack.isOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.push(int(val))
					keepNum = ""
				}
			}
			// val, _ := strconv.ParseInt(ch, 10, 64)
			// numStack.push(int(val))

		}

		//exp是否掃描完了

		if index+1 == len(exp) {
			break
		}
		index++
	}

	//4. 如果整個exp 都遍歷完，就依次將 operStack 取1個運算符 、numStack取2個數 來運算，最後把結果放入numStack，直到operStack為空
	for {
		if operStack.top == -1 {
			break
		}
		num1, _ = numStack.pop()
		num2, _ = numStack.pop()
		oper, _ = operStack.pop()
		result = operStack.cal(num1, num2, oper)

		//將最終計算結果放入numStack
		numStack.push(result)

	}

	res, _ := numStack.pop()
	fmt.Printf("表達式%s = %v", exp, res)

}

type stack struct {
	maxTop int // 棧頂最大值
	top    int // 棧頂
	//可以不用寫棧底，因為其是固定的 0 or -1

	arr [20]int //陣列
}

//判斷字符是不是運算符[+-*/]
//用acsii碼來判斷
func (sta *stack) isOper(val int) bool {

	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	}

	return false

}

//預算的方法
func (sta *stack) cal(num1 int, num2 int, oper int) int {
	res := 0
	//stack 先進後出，但我們計算都是從前面往後算的，num2 、num1 位置要放對
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("oper err")
	}

	return res
}

//返回某個運算符的優先級 (可自訂義)
func (sta *stack) prior(oper int) int {
	res := 0
	//*/ 返回1
	//+- 返回0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
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

//棧 運用在表達式  ex: 3+3*6-2

//1. 創建兩個stack  (數棧)numStack、(運算符棧)operStack

//2. 將表達式存在一個字串 exp  ，並遍歷他，如果是數字 => numStack

//3. 如果是運算符
//  (1)如果operStack 是空stack => 把她加入operStack
//         處理多位數 =>主要是處理掃描到數的時候，要再往index後探一位
//         a. 如果到尾端，就不往後了
//         b. 如果index後一位是 "運算符" 把數加入numStack吧
//         c. 把數拼接起來 ，再加入numStack
//  (2)如果不是 空stack
//     a. 如果operStack "頂端的" 符號  ， 優先運算  於  "準備加入的" 符號，  (先*/後+-)
//        從 operStack 取出 符號 ， numStack 也取出兩個數 來運算，
//        運算完把結果重新放到 numStack  ，"準備加入的" 運算符 ，再加入operStack

//     b. 如果operStack "準備加入的" 符號 ， 優先運算  於  "頂端的" 符號 => 運算符 直接加入operStack

//4. 如果整個exp 都遍歷完，就依次將 operStack 取1個運算符 、numStack取2個數 來運算，最後把結果放入numStack，直到operStack為空
