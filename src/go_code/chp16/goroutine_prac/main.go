package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//練習
	//主線程 每隔一秒打印hello world 輸出10後退出
	//協程   每隔一秒打印hello go 輸出10後退出
	go test() //加關鍵字 go 就變協程了 (主線程執行完，就算協程還沒執行完也會被停止)
	for i := 1; i <= 5; i++ {
		fmt.Println("hello world", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test() {
	for i := 1; i <= 7; i++ {
		fmt.Println("hello go", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

//goroutine

// go的 進程、線程(併行、併發)
// 1. 進程(process): 程序在操作系統的一次執行過程， 是操作系統進行資源分配的基本單位
// 2. 線程(thread): 程序執行的最小單位， 線程是進程的一個執行實例(執行緒)

// ex: 開啟Foxy => 一個進程  ，  裡面可以載很多A片=> 每個下載過程都是一個 線程

// 3. process 可以創立或銷毀多個thread ，process中可以有多個thread 併發執行
// 4. 程序至少有一個process，一個process 至少有一個 thread

// go 的併發、併行
// 1. 併發: 多個thread 在單核執行 => 傳統語言的多線程是這樣操作
// 2. 併行: 多個thread 在多核執行 =>比併發還快

// (但web的多併發真的是同時)

// 併發: 外在角度看是同時執行，但從微觀角度上，一個時間只有一個線程，只是切換(輪詢操作)很快
// 併行: 在多個CPU上真的是同時執行，微觀角度也是

// go的 主線程 、 協程
// 1. 主線程: 當作傳統process吧
// 2. 協程: thread =>輕量級的線程

// 主線程是一個物理線程，直接作用在CPU上的。是重量級，非常耗費CPU資源。 =>如果主線程結束，協程就算還沒執行完也會結束
// 協程是從主線程開啟的，是輕量級的線程，是邏輯態。對資源消耗相對小

// go的協程機制是重要的特點，可以開啟上萬個協程。其他邊程語言的併發機制是一般基於線程的，開啟過多的線程，資源耗費大，這裡就凸顯GO在併發上的優勢了

// 一個go 主線程，可以有多個協程(goroutine)
// goroutine特性
// a. 有獨立的棧空間 (stack)
// b. 共享堆空間 (heap)  =>引用類型存放
// c. 調用由用戶控制
// d. 輕量級的線程

//MPG模式(goroutine的調度模型)
//M:主線程
//P:協程執行需要的上下文(資源)，可以看程調度器
//G:協程

//如果多個M都在同個CPU執行: 併發
//多個M在不同CPU執行 : 併行

//運作模式: 有個M0(主線程)  ， 在執行一堆G(協程) 的列隊， P 負責把G放到G0執行，當 G0在處理資料庫，遇到阻塞，其他的G 正在等待中
//P 會把其他的G 列隊調度到M1(新增一個主線程)去執行，確保G0在執行時，其他列隊的G不會被阻塞
//等到G0執行完了，M會被放到空閒的主線程，G0繼續執行下一個協程