package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	//定義一般的陣列/切片
	var intSlice = []int{0, -1, 10, 7, 20}
	fmt.Println(intSlice)
	//排序切片 sort包  func Ints(a []int)
	//1. 冒泡排序
	//2. 官方內建函數

	//打sort 快捷鍵會冒出下列
	// type SortBy []Type

	// func (a SortBy) Len() int           { return len(a) }
	// func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
	// func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }
	fmt.Println("排序後----")
	sort.Ints(intSlice) //不用&intSlice，切片本身是引用類型
	fmt.Println(intSlice)

	//對一個struct 切片排序
	//sort包  func Sort(data Interface)   可以傳入實現Interface的變數
	var heros HeorSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//將hero append 到切片
		heros = append(heros, hero)
	}

	//排序前
	for _, val := range heros {
		fmt.Println(val)
	}

	//排序後
	sort.Sort(heros)
	fmt.Println("排序後----")
	for _, val := range heros {
		fmt.Println(val)
	}

}

//定義一個結構體

//Hero is a struct
type Hero struct {
	Name string
	Age  int
}

//定義一個struct 切片 來裝Hero

//HeorSlice is a slice for Hero struct
type HeorSlice []Hero

//要使用sort包 func Sort(data Interface) ，就要實現Interface 介面裡所有方法

//切片長度
func (hs HeorSlice) Len() int {
	return len(hs)
}

//決定要用何種標準排序
func (hs HeorSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age //小到大排序(年齡)  可以看自己要以什麼屬性排序
}

//互換元素
func (hs HeorSlice) Swap(i, j int) {

	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	//也可以這樣寫
	hs[i], hs[j] = hs[j], hs[i]
}

//(接口)介面編程
//實現可以排序struct 切片 sort.Sort(data Interface)

//sort.Sort(介面)
//=> 如果要實作，就是要完全實作介面裡所有方法
//=> 定義一個struct 裡面實作 len()、Less()、Swap()
//就可以調用sort.Sort(介面)
