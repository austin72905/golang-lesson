package main

import (
	"fmt"
	"sort"
)

func main() {

	intMap := map[int]string{
		1: "jan",
		2: "feb",
		3: "mar",
		4: "apr",
		5: "may",
	}
	var saveKeys []int
	for keys := range intMap {
		//將 map的key裝入slice
		saveKeys = append(saveKeys, keys)
	}
	//sort 包裡面的方法，對slice裡元素進行排序，返回一個slice
	sort.Ints(saveKeys)
	fmt.Println(saveKeys)

	//逐一輸出
	for _, keys := range saveKeys {
		fmt.Printf("map[%v] = %v\n", keys, intMap[keys])
	}
}

//map 排序

//1. 先用切片 裝入map的key 然後排序 ，之後再依序輸出
