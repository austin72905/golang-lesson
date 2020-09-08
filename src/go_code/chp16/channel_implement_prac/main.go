package main

import (
	"fmt"
	"time"
)

func main() {
	//思路
	//1. 兩個協程都是操作同一個channel
	//2. 建立兩個channel ， 一個裝整數(chan a)，一個裝bool(chan b)
	//3. 當readData 讀取完所有數據，便向chan b 寫入true，並關閉channnel
	//4.主線程可以用一個for迴圈，不斷向chan b 讀取，當讀到true (代表協程都工作完了)，跳出程序

	//建立兩個channel
	intChan := make(chan int, 50)
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
		time.Sleep(time.Millisecond * 50)
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
		fmt.Println("讀到數據 =", val, " ,readDown =", readDown)
	}
	//讀完全部的數據
	downChan <- true
	close(downChan)
}

//應用

//需求
//1. 開啟一個writeData 協程 向channel intChan 寫入50個數
//2. 開啟一個readData  協程 從channel intChan 讀取數據
//3. 主線程等協程都執行完畢才退出
