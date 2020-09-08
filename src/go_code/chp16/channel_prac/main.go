package main

import (
	"fmt"
	"sync"
	"time"
)

//練習
//計算1~20的乘階，並把各個數的乘階n!放入map

//1. 寫一個函數，計算各個數的乘階，並放入map
//2. 運用到協程，map要全局的
//3. 思考一個問題 寫進map能不能用併發? : 答，不行，會爆炸(讀可以)

//如果這樣寫會報錯  fatal error: concurrent map writes(併發寫入)

//全局變數
// var (
// 	myMap = make(map[int]int, 10)
// )

// func testGo(n int) {
// 	result := 1
// 	for i := 1; i <= n; i++ {
// 		//計算階乘
// 		result *= i
// 	}
// 	//將結果寫入map
// 	myMap[n] = result
// }

// func testGo(n int) {
// 	result := 1
// 	for i := 1; i <= n; i++ {
// 		//計算階乘
// 		result *= i
// 	}
// 	//將結果寫入map

// 	myMap[n] = result

// }

// func main() {

// 	for i := 1; i <= 20; i++ {
// 		go testGo(i)
// 	}

// 	//讓主線程休眠5秒，目的讓協程完成任務
// 	time.Sleep(time.Second * 5)

// 	//打印map內容
// 	for index, val := range myMap {
// 		fmt.Printf("myMap[%d]=%d\n", index, val)
// 	}

// }

//解決: 全局變數加鎖，如果協程要處理這個變數，會先把全局變數鎖上(lock)，等處理完再解鎖(unlock)，
//當其他協程來時，發現鎖被鎖上，就要等待(排隊)

//解決方法
//全局變數
var (
	myMap = make(map[int]int, 10)
	//(解決)新增一個全局互斥鎖
	lock sync.Mutex
)

func testGo(n int) {
	result := 1
	for i := 1; i <= n; i++ {
		//計算階乘
		result *= i
	}
	//將結果寫入map
	//加鎖
	lock.Lock()
	myMap[n] = result
	//解鎖
	lock.Unlock()
}

func main() {
	//練習
	//計算1~200的乘階，並把各個數的乘階n!放入map

	//1. 寫一個函數，計算各個數的乘階，並放入map
	//2. 運用到協程，map要全局的
	//3. 思考一個問題 寫進map能不能用併發? : 答，不行，會爆炸(讀可以)

	//如果這樣寫會報錯  fatal error: concurrent map writes(併發寫入)
	//解決: 全局變數加鎖，如果協程要處理這個變數，會先把全局變數鎖上(lock)，等處理完再解鎖(unlock)，
	//當其他協程來時，發現鎖被鎖上，就要等待(排隊)
	for i := 1; i <= 20; i++ {
		go testGo(i)
	}

	//讓主線程休眠5秒，目的讓協程完成任務
	time.Sleep(time.Second * 5)

	//照理說5秒上面的協程都應該執行完了，後面就不會出現資源競爭的問題，但在實際運行中，
	//我程序設計上知道5秒就執行完所有協程，但主線程不知道，底層可能仍會出現資源競爭，所以還是要加鎖

	lock.Lock()

	//打印map內容
	for index, val := range myMap {
		fmt.Printf("myMap[%d]=%d\n", index, val)
	}

	lock.Unlock()

}

//goroutine 之間的通訊 channel(管道)

//解決主線程跑完，協程還沒跑完

//資源競爭 編譯程序時 加 -race  => go build -race {檔名}.go  =>數據有競爭關係會幫你打印出來

//sync (同步)包 的 func (m *Mutex) Lock()      Mutex=>互斥鎖
