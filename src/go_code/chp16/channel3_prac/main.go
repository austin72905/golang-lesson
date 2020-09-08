package main

import "fmt"

func main() {

	//var allChan chan interface{}
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "austin"
	cat := Cat{"miao", 3}
	allChan <- cat

	//希望獲得第3個元素，要先把前兩個推掉
	<-allChan
	<-allChan
	newCat := <-allChan                                     //從管道取出來的貓
	fmt.Printf("newCat = %T,newCat = %v\n", newCat, newCat) //newCat = main.Cat,newCat = {miao 3}

	//神奇的事發生了，下面會編譯不過
	//因為本質上newCat 還是一個空接口，上面是因為在運行的瞬間，發現是什麼類型，下面是編譯的時候就發現錯誤了
	//fmt.Printf("newCat.Name=%v", newCat.Name) //newCat.Name undefined (type interface {} is interface with no methods)

	//解決: 類型推斷
	newCatAssert := newCat.(Cat) //看 newCat 是否能被轉為 Cat struct
	fmt.Printf("newCat.Name=%v", newCatAssert.Name)
}

//Cat struct
type Cat struct {
	Name string
	Age  int
}
