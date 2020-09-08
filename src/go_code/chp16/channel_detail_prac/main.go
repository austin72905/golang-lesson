package main

import (
	"fmt"
	"time"
)

func main() {
	//聲明只能寫
	var onlyReadChan chan<- int
	onlyReadChan = make(chan int, 3)
	onlyReadChan <- 20

	//聲明只能讀
	// var onlyWriteChan <-chan int
	// onlyWriteChan = make(chan int, 3)
	// //onlyWriteChan <- 20   無法寫進去， 會報錯 invalid operation: onlyWriteChan <- 20 (send to receive-only type <-chan int)
	// <-onlyWriteChan

	//實際開發中，有時候不確定何時可以關閉channel，
	//可運用select

	intChan := make(chan int, 10)
	for i := 0; i < 5; i++ {
		intChan <- i
	}

	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "no" + fmt.Sprintf("%d", i)
	}

	//select 能在不關閉channel的情況下去歷遍，當取不到值，他會跑下一個case
	for {
		select {
		case val := <-intChan:
			fmt.Printf("從intChan取出數據%d\n", val)
			time.Sleep(time.Second)
		case val := <-strChan:
			fmt.Printf("從intChan取出數據%s\n", val)
			time.Sleep(time.Second)
		default:
			fmt.Printf("取不到數據了\n")
			return
		}
	}
}

//channel 細節

//1. channel 可聲明為 只讀 ，或是 只寫

//2. 默認是雙向的

//3. 只讀只寫可以應用在函數的"參數"  ，讓函數在操作該管道時，不會誤操作

//4. select : 解決 channel 阻塞問題  (傳統方法要關閉channel才能遍歷，不然會deadlock)
