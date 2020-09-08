package main

import (
	"errors"
	"fmt"
)

func main() {
	//正常執行到錯誤下面就不會在執行了
	//testErr()

	//但可以用錯誤處理讓程式繼續執行
	testErr1()
	fmt.Println("main().....")
	fmt.Println("--------自定義錯誤練習-------------")
	customErr()
	fmt.Println("成功執行")

}

//錯誤的函數
func testErr() {
	num1 := 10
	num2 := 0
	//這邊會報錯因為 除數為0
	res := num1 / num2
	fmt.Println(res)
}

//處理方法
func testErr1() {
	//defer 一個匿名函數，在testErr整個執行完才會跑他 ， 呼叫匿名函數方法 在後面加()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("錯誤發生 :", err)
		}
	}()
	num1 := 10
	num2 := 0
	//這邊會報錯因為 除數為0，但有defer後在此函數報錯之後的部分不會執行，會直接執行defer，函數跳出後不會崩潰，之後的代碼會繼續執行
	res := num1 / num2
	fmt.Println(res)
}

//可以自定義錯誤練習
//假設使用函數去讀取config 訊息，如果讀取文件名不正確，返回一個自定義的錯誤
func readConfig(fileName string) (err error) {
	//如果輸入的文件名正確，返回空的error值
	if fileName == "config.init" {
		return nil
	}
	return errors.New("讀取配置....失敗")

}

func customErr() {
	fileName := "config.ini"
	err := readConfig(fileName)
	//如果回傳的error不是空值，讓panic終止程序
	if err != nil {
		panic(err)
	}

	fmt.Println("後面的代碼繼續執行")
}

//錯誤處理

//默認發生錯誤(panic)，程序會崩潰退出

//go 不支持 try...catch...finally

//go 的錯誤處理為 defer , panic, recover

// go 可以拋出一個panic 異常 ，在defer中 用recover 來catch這個異常

//可以自定義錯誤
//用err包裡 func New(text string) error  , 返回一個error類型
//內建函數 panic 接收interface{} 類型的值(也就是任何值) 作為參數，可以接受error 類型的變數，輸出錯誤訊息，並退出程序
