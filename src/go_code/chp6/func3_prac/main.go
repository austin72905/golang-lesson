package main

import (
	"fmt"
	"strings"
)

func main() {

	//f 現在是一個函數可以呼叫
	f := AddUpper()
	fmt.Println(f(1)) //11  n的植被改變了(n變成了11)
	fmt.Println(f(2)) //13

	f2 := makeSuffix(".jpg")
	fmt.Println("文件名=", f2("win"))
	fmt.Println("文件名=", f2("winter.jpg")) //如果這個字串有jpg就不會幫他加了
}

//AddUpper return  func
//累加器
func AddUpper() func(int) int {
	//這邊會形成閉包 n的值每呼叫一次不斷的往上加，下次調用不會初始化，而是累加
	var n int = 10
	//返回的是一個函數，但這個函數引用了函數外的變數n，所以此函數與 n 形成了閉包
	return func(x int) int {
		n = n + x
		return n
	}
}

//練習
// 1. 寫一個函數 makeSuffix(suffix string) 可以接收一個後墜名(ex.jpg)，並返回一個函數
// 2. 調用makeSuffix 可以傳入一個字串，如果沒有指定後綴，則返回字串.jpg
// 3. 要以閉包的方式
// 4. 提示: strings.HasSuffix ,可以判斷某字符串是否有後綴
func makeSuffix(suffix string) func(string) string {

	//上方傳入的suffix 與下面返回的函數會形成閉包，suffix的值會被保存
	return func(name string) string {
		//func HasSuffix(s, suffix string) bool
		//判断s是否有后缀字符串suffix
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

//閉包
// 一個函數 與 其相關的引用變數 組合成的一個整體
//要分析出返回的函數使用到那些變數
//不用閉包也是可以完成功能，但用了你參數只要寫一次，可重複利用
