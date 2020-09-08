package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHi()
	go unmake()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}

}

func sayHi() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hi")
	}
}

//發生錯誤的協程
func unmake() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("func unmake occured err errMsg:", err)
		}
	}()

	//未 make 即使用map
	var myMap map[int]string
	myMap[0] = "go"
}

//使用defer+recover ，當一個協程panic時，不會整個程序崩潰
