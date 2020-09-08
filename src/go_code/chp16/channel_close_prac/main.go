package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	//關閉channel
	close(intChan)
	//關閉之後不能寫數據進去，但能讀
	//intChan <- 200   //panic: send on closed channel
	n1 := <-intChan
	fmt.Println(n1)

	//遍歷
	intChan2 := make(chan int, 20)

	//寫入數據 用 for 可以
	for i := 0; i < 20; i++ {
		intChan2 <- i * 2
	}

	//遍歷channel要先關閉，不然會報fatal error: all goroutines are asleep - deadlock!
	close(intChan2)

	for val := range intChan2 {
		fmt.Println("val =", val)
	}

}

//channel的關閉

//內建函數builtin 包的 func close(c chan Type)

//關閉之後不能再向channel寫數據，但可以讀

//channel的遍歷 (for range)  也是可以用for ，但因為長度會動態變化，取出一個channel長度就-1，要注意

//channel 用for range 沒有index，只有val，因為他是queue
//如果沒有關閉channel會報deadlock
