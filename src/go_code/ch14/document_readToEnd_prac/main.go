package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	//打開文件
	file := "C:/Users/USER/Desktop/文件檔/test.txt"
	content, err := ioutil.ReadFile(file) // Open 、Close都被封裝在ReadFile裡面了
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}

	//將讀取的內容顯示到終端
	fmt.Printf("%v", string(content)) // []byte 要用string()轉，不然是一堆數字

}

//一次性讀取文件(適合較小的) ioutil 包  func ReadFile(filename string) ([]byte, error)
