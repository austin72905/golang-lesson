package main

import "fmt"

func main() {
	//但可以用for 實現
	var i = 1
	for {
		if i > 10 {
			break
		}
		fmt.Println("hello")
		i++
	}
	//do while
	var j = 1
	for {
		//一定會先執行一次
		fmt.Println("hello")
		j++
		if j > 10 {
			break
		}

	}
}

//golang 裡沒有while 跟 do while，但可以用for 實現
