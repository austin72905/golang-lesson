package main

import (
	"fmt"
	//也可以打絕對路徑C:\Users\USER\GoProj\src\go_code\chp3\pack_prac\bag
	//但我們設定了GOPATH ，系統默認會加src
	"go_code/chp3/pack_prac/bag"
)

func main() {
	//大寫的變數或涵式才可被使用
	fmt.Println(bag.Her)
}
