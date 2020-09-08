package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "C:/Users/USER/Desktop/文件檔/test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("err occured when open errMsg=%v", err)
		return
	}

	//記得關閉
	defer file.Close()

	//帶緩存的讀法
	reader := bufio.NewReader(file)

	//new 一個 struct
	var charCount charCount //預設裡面屬性int 值都是0
	//開始一行一行的讀
	for {
		//帶緩存的讀法，讀一行，處理一行，
		content, err := reader.ReadString('\n') //一行一行的讀 ，這個方法結尾一定要再換行，不燃讀不到
		if err == io.EOF {
			break //讀到底就退出
		}
		//讀出來的內容是ascii  A~Z、a~z、0~9、其他符號
		//轉成這個格式才能統計中文
		//contents := []rune(content)
		//進行統計

		for _, val := range content {

			//fmt.Println(val)
			//fallthrough 穿透，會繼續執行下一個case裡面的內容
			switch {
			case val >= 'a' && val <= 'z':
				fallthrough
			case val >= 'A' && val <= 'Z':
				charCount.eng++
			case val == ' ' || val == '\t':
				charCount.space++
			case val >= '0' && val <= '9':
				charCount.num++
			default:
				charCount.other++
			}
		}
	}

	fmt.Printf("字符數量=%v  數字數量=%v  空格數量=%v  其他字符數量=%v\n", charCount.eng, charCount.num, charCount.space, charCount.other)

}

//用strcut 來裝統計數字
type charCount struct {
	num   int
	eng   int
	space int
	other int
}

//練習統計一個文件裡有多少中文、英文、數字、空格、其他字符

//1. 打開文件，創建一個reader

//2. 一行一行的讀

//3. 用strcut 保存
