package main

import "fmt"

func main() {
	//練習
	//使用map紀錄monster 訊息 name、age，一個monster對應一個map
	monster := make([]map[string]string, 1) //monster is a slice of map   size不可省略

	if monster[0] == nil {
		monster[0] = map[string]string{
			"name": "snake",
			"age":  "200",
		}
	}
	fmt.Println(monster)
	//如果要動態增加可用 append  => 1. 先定義一個新map 2. 把他加入map 切片
	//1. 先定義一個新map
	newMonster := map[string]string{
		"name": "dragon",
		"age":  "1000",
	}
	//2. 把他加入map 切片
	monster = append(monster, newMonster)
	fmt.Println(monster)

}

//map 切片

//目的:可以動態增加map
