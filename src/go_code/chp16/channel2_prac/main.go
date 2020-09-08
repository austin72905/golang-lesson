package main

import "fmt"

func main() {
	//聲明一個channel
	var intChan chan int
	intChan = make(chan int, 3) // 3=>開一個容量3給他

	//channel是引用類型，本身存放一個地址
	fmt.Printf("intChan本身的值=%v  intChan所在地址=%v\n", intChan, &intChan)

	//寫入數值到channel
	intChan <- 10 //可以寫常數進去，也可以自定義變數寫進去
	num1 := 12
	intChan <- num1

	fmt.Printf("intChan 的長度=%v 容量=%v\n", len(intChan), cap(intChan)) //len=2 cap=3

	intChan <- 11
	//現在intChan [10,12,11]
	//寫入的數據超過channel的容量會爆炸
	//intChan <- 13
	fmt.Printf("intChan 的長度=%v 容量=%v\n", len(intChan), cap(intChan)) //fatal error: all goroutines are asleep - deadlock!

	//解決: 取出數據過後容量空出來就可以繼續放了
	//如果取出的數據不用了 可以 <-intChan  (不指定變數)
	num2 := <-intChan
	fmt.Println("num2 =", num2)                                      //10  先進先出
	fmt.Printf("intChan 的長度=%v 容量=%v\n", len(intChan), cap(intChan)) //len=2 cap=3

	num3 := <-intChan //12
	num4 := <-intChan //11

	fmt.Println("num3 =", num3, ",num4 =", num4)
	fmt.Printf("intChan 的長度=%v 容量=%v\n", len(intChan), cap(intChan)) //len=0 cap=3

	//現在intChan [] 全部取完了
	//在沒有使用協程情況下，如果channel數據已經全部取出，再取就會報deadlock
	//num5 := <-intChan //
	//fmt.Println("num5 =", num5)//fatal error: all goroutines are asleep - deadlock!
}

//channel 說明

//為什麼需要?
//全局變數+ lock 解決goroutine不完美
//1. 主線程等在所有G 全部完成的時間難以確定，設定秒數僅是估算
//2. 如果主線程休眠時間長，等待時間會浪費，如果短了G可能還未工作完而被銷毀
//3. 全局變數+lock不利於多個協程對全局變數的讀寫操作

//基本介紹
//1. channel 本質是 queue
//2. 先進先出
//3. 線程安全，多個G訪問時，不需要lock，不會有資源競爭問題
//4. channel有類型，string 的 channel 只能放string  (如果要放多種數據可用interface{}，但不建議)

//聲明方法
//1. var 變數 chan 數據類型   =>(只能存該數據類型)
//2. channel 是引用類型
//3. 必須先make才能用
//4. 容量(cap)不會變化，寫入channel超過容量會爆炸
//5. 在沒有使用協程情況下，如果channel數據已經全部取出，再取就會報deadlock
