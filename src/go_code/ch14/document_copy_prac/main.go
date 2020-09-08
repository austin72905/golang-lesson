package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//將一張圖片/電影/MP3 複製到另一個文件
	srcFile := "C:/Users/USER/Desktop/myIceCream.jpg" //png 好像不能寫
	dstFile := "C:/Users/USER/Desktop/文件檔/myIceCream.jpg"
	_, err := copyFile(dstFile, srcFile)
	if err != nil {
		fmt.Printf("err occured when copy errMsg:%v\n", err)
	} else {
		fmt.Println("copy succeed")
	}
}

func copyFile(dstFileName string, srcFileName string) (written int64, err error) {

	//讀取文件
	srcFile, err := os.Open(srcFileName)

	if err != nil {
		fmt.Printf("err occured when openging... errMsg:%v\n", err)
		return
	}
	// 結束後關閉檔案
	defer srcFile.Close()
	// 源文件的reader
	srcReader := bufio.NewReader(srcFile)

	//打開要寫入的目的路徑，如果不存在就建一個
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666) //0666 windows 無效果

	if err != nil {
		fmt.Printf("err occured when opening... errMsg:%v\n", err)
		return
	}
	//結束記得關閉
	defer dstFile.Close()

	//目標文件的 writer
	dstWriter := bufio.NewWriter(dstFile)

	//複製文件
	return io.Copy(dstWriter, srcReader)

}

//copy 文件

// io 包的  func Copy(dst Writer, src Reader) (written int64, err error)

// dst 目標路徑 的writter  src  來源路徑的reader
