package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//獲取當前時間
	t := time.Now()
	fmt.Printf("now = %v , type = %T\n", t, t)
	//now = 2020-08-08 19:53:56.1379267 +0800 CST m=+0.002980001 , type = time.Time
	fmt.Println("獲取  -----年月日、時分秒-----")
	//獲取年月日、時分秒
	fmt.Printf("年=%v\n", t.Year())
	fmt.Printf("月=%v\n", t.Month()) //數字: int(t.Month())
	fmt.Printf("日=%v\n", t.Day())
	fmt.Printf("時=%v\n", t.Hour())
	fmt.Printf("分=%v\n", t.Minute())
	fmt.Printf("秒=%v\n", t.Second())

	//格式化日期、時間
	//1. 方法一
	fmt.Printf("當前時間 %d-%d-%d %d:%d:%d \n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	//也可以把印出的內容存成字串  fmt.Sprintf()
	dtStr := fmt.Sprintf("當前時間 %d-%d-%d %d:%d:%d \n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	fmt.Printf("dtStr=%v\n", dtStr)
	//2. 方法二 Time.Format()
	//2006/01/02 15:04:05 這個值是固定的，據說是作者想做go語言的時間XD

	//如果要輸出到秒
	//                    年  月 日  時 分 秒
	fmt.Printf(t.Format("2006/01/02 15:04:05"))
	fmt.Println()
	//只要輸出年月日
	fmt.Printf(t.Format("2006/01/02"))
	fmt.Println()
	//輸出時間
	fmt.Printf(t.Format("15:04:05"))
	fmt.Println()
	//可以自由組合
	fmt.Printf(t.Format("2006"))
	fmt.Println()

	//時間常數
	// type Duration int64 //自定義的type

	// const (
	// 	Nanosecond  Duration = 1                  //納秒
	// 	Microsecond          = 1000 * Nanosecond  //微秒
	// 	Millisecond          = 1000 * Microsecond //毫秒
	// 	Second               = 1000 * Millisecond //秒
	// 	Minute               = 60 * Second        //分
	// 	Hour                 = 60 * Minute        //時
	// )
	//如果要用0.1秒 Second*0.1 =>xx 不能這樣用，要 100*Millisecond

	//睡眠
	// for i := 0; i < 30; i++ {
	// 	fmt.Println(i)
	// 	time.Sleep(time.Millisecond * 100) //0.1秒

	// }

	//unix(秒) 、unixnano(納秒) 時間戳 使用
	fmt.Printf("unix時間戳=%v  ,  unixnano時間戳=%v\n", t.Unix(), t.UnixNano())

	//練習
	//計算涵式執行時間

	test()

}

func test() {
	start := time.Now().Unix()
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
	end := time.Now().Unix()
	fmt.Printf("涵式執行耗費時間= %v 秒\n", end-start)

}

//時間函數

//1. 獲取當前時間 func Now() Time  ,Time 是一種資料型別

//2. 獲取年月日、時分秒

//3. 格式化日期、時間

//4. 時間的常數   time包裡面 有定義了一堆時間單位的常數，

//5. 睡眠，func Sleep(d Duration)  過幾秒執行動作

//6. unix(秒) 、unixnano(納秒) 時間戳 (可以配合random獲取隨機數字)  func (t Time) Unix() int64  ,1970.1.1到指定時間所過的秒數
