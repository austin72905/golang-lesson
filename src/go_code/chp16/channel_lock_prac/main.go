package main

import (
	"fmt"
	"time"
)

func main() {

	//建立兩個channel
	intChan := make(chan int, 10)
	downChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, downChan)

	//阻塞主線程，要等到協程都工作完才跳出程序
	for {
		_, allWorkDowm := <-downChan
		if allWorkDowm == false {
			break
		}
	}
}

//write data
func writeData(intChan chan int) {
	//寫入數據
	for i := 1; i <= 50; i++ {
		intChan <- i
		//time.Sleep(time.Millisecond * 500)
		fmt.Println("寫入數據 = ", i)
	}
	//關閉intChan
	//為了方便用for 才用close
	close(intChan)
}

//read data
func readData(intChan chan int, downChan chan bool) {

	for {
		val, readDown := <-intChan
		//close 會讓被關閉的channel回傳兩個值，當取出所有數據時，會把第二個值設為false
		if readDown == false {
			break
		}
		time.Sleep(time.Millisecond * 500)
		fmt.Println("讀到數據 =", val, " ,readDown =", readDown)
	}
	//讀完全部的數據
	downChan <- true
	close(downChan)
}

//channel 阻塞

//如果一個channel 只有寫沒有讀，channel會阻塞，最後報deadlock

//但測試如果在主線乘用time.Sleep()可以規避

//只要有讀，就算讀得很慢，不會deadlock，只是channel會形成有效阻塞(就算容量10，而寫入50個數據)
