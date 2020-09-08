package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	//打開文件
	file, err := os.Open("C:/Users/USER/Desktop/文件檔/test.txt")
	if err != nil {
		fmt.Println("open file err =", err)
	}

	//關閉文件
	defer file.Close()

	//創建一個*Reader ，是帶緩衝的(不是一次全讀取，讀取一部份，處理一部份)
	reader := bufio.NewReader(file)

	//循環讀取文件內容
	for {
		str, err := reader.ReadString('\n') //讀到換行就結束
		if err == io.EOF {                  //io.EOF表示文件末尾
			break
		}
		//輸出內容
		fmt.Print(str) //不要Println 不然會有兩個換行符
	}

	fmt.Println("read to end...")

}

//讀取文件內容顯示在終端(帶緩衝區)  func NewReader(rd io.Reader) *Reader   =>Reader 點進去看源碼

// Buffered input.
//預設緩衝4096字節
// const (
// 	defaultBufSize = 4096
// )
