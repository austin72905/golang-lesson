package main

import (
	"fmt"
	"time"
)

func main() {

	//建立一個管道intChan 來裝1~8000
	intChan := make(chan int, 1000)

	//判斷是質數的數 放入 管道primeChan
	primeChan := make(chan int, 3000)

	//exitChan
	exitChan := make(chan bool, 4)

	//寫入 1~8000 到管道intChan
	go putNum(intChan)

	//建立4個協程 primeNum1~4 從 intChan 取數據出來判斷哪些是質數
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//主線程阻塞，等待所有協程完成才跳出程序
	go func() {
		for i := 0; i < 4; i++ {
			//當取不出東西的時候，就會卡在這邊，不會進入下一個i
			<-exitChan
		}

		//當exitChan 取出4個true，代表協程都執行完了，可以關閉
		close(primeChan)
	}()

	//遍歷結果
	for {
		val, read := <-primeChan
		if read == false {
			break
		}
		//將結果輸出
		fmt.Printf("質數=%d\n", val)
	}

	fmt.Print("程序結束")

}

//寫入 1~8000 到管道intChan
func putNum(intChan chan int) {
	//質數從2開始
	for i := 2; i <= 8000; i++ {
		intChan <- i
	}

	//寫完就可以關了
	close(intChan)
}

//建立4個協程 primeNum1~4 從 intChan 取數據出來判斷哪些是質數
//其實可以建立一個，用for迴圈跑4次就好
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	//是否是質數
	var isPrime bool

	//var isPrime bool=true  不要這樣寫這樣isPrime 在新的迴圈不會重新變true

	for {
		time.Sleep(time.Millisecond * 10)
		//有close了當讀不到數據，readDone會返回false
		val, read := <-intChan
		if read == false {
			//當讀不到東西時跳出迴圈
			break
		}
		isPrime = true
		//質數?  (n (從2開始) 除以1~n-1 ， 如果都不能被整除就是質數)
		for i := 2; i < val; i++ {
			//可以被其他數整除代表不是質數
			if val%i == 0 {
				isPrime = false
				break
			}
		}

		//將質數丟入管道 primeChan
		if isPrime {
			primeChan <- val
		}
		//這邊不能關閉，因為其他協程可能還要寫數據進去
	}

	fmt.Println("有個協程讀不到數據而閒置了...")

	//向 exitChan 寫入ture，表示已完成
	exitChan <- true
}

//練習: 統計1~8000 哪些是質數?  (n (從2開始) 除以1~n-1 ， 如果都不能被整除就是質數)
//思路
// 1. 建立一個管道intChan 來裝1~8000
// 2. 建立4個協程 primeNum1~4 從 intChan 取數據出來判斷哪些是質數
// 3. 協程 primeNum1~4 判斷是質數的數 放入 管道primeChan
// 4. 協程 primeNum1~4 工作完之後，分別向管道 exitChan 寫入true  (協程 primeNum 可以不用等到其他primeNum都工作完才停止)
// 5. 主線程使用for迴圈 ， 不斷向管道 exitChan 讀數據 => 如果讀到4個true，程序跳出
