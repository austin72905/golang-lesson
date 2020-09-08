package main

import (
	"fmt"
	"os"
)

func main() {
	//file 別名
	//1. 文件對象
	//2. 文件指針
	//3. 文件句炳

	//打開文件
	file, err := os.Open("C:/Users/USER/Desktop/文件檔/test.txt")
	if err != nil {
		fmt.Println("open file err =", err)
	}

	//輸出文件
	fmt.Printf("file=%v\n", file) //輸出一個地址 (所以文件是一個指針)

	//關閉文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err =", err)
	}
}

//文件

//是以數據流(數據源(文件)與內存(程式)之間經歷的路徑)來操作

//輸入流: 讀文件

//輸出流: 寫文件

//go 對文件的操作是用 os 包  裡面有個File struct  (文件是一個指針類型  f*File)

//基本操作

//打開文件  func Open(name string) (file *File, err error)

//關閉文件  func (f *File) Close() error

//讀取文件內容顯示在終端(帶緩衝區)
