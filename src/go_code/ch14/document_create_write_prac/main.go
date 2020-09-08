package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	//主要在更換flag

	//練習一， 創立一個新文件，寫入5個 hello,world

	//寫入: bufio.NewWriter()  帶緩存的寫入
	//createAndWrite()

	//練習二， 打開一個存在的文件，將原來的內容覆蓋成10句 襙你媽的
	//coverAndWrite()

	//練習三， 在原來的內容追加 haha, catch you
	//appendAndWrite()

	//練習四， 打開一個存在的文件，將內容顯示在console，追加5句taiwan no.1
	//readAndAppend()

	//練習五， 將一個文件的內容寫入另一個文件(兩個文件已存在)
	//使用ioutil
	//1. 讀取文件內容到緩存
	//2. 寫入
	writeToOther()
}

//創立一個新文件，寫入5個 hello,world
func createAndWrite() {
	fileName := "C:/Users/USER/Desktop/文件檔/createAndWrite.txt"

	//開啟文件的模式可用 | 混用    0666 =>windows 沒用
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("err occured when open file.. msg :%v\n", err)
		return
	}
	//關閉檔案
	defer file.Close()

	wrtContent := "hello,world\r\n" //如果只寫\n 記事本認不出來
	//寫入時是帶緩存的
	writer := bufio.NewWriter(file) //func NewWriter(w io.Writer) *Writer
	for i := 0; i < 5; i++ {

		writer.WriteString(wrtContent)

	}

	//writer 是帶緩存的，其實是先寫到混存裡面，要Flush才真正寫入文件
	writer.Flush()
	fmt.Println("write succeed")

}

//打開一個存在的文件，將原來的內容覆蓋成10句 襙你媽的
func coverAndWrite() {
	fileName := "C:/Users/USER/Desktop/文件檔/createAndWrite.txt"

	//開啟文件的模式可用 | 混用    0666 =>windows 沒用
	//os.O_TRUNC 會把文件清掉
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0666)

	if err != nil {
		fmt.Printf("err occured when open file.. msg :%v\n", err)
		return
	}
	//關閉檔案
	defer file.Close()

	wrtContent := "haha, catch you\r\n"
	//寫入時是帶緩存的
	writer := bufio.NewWriter(file) //func NewWriter(w io.Writer) *Writer
	for i := 0; i < 5; i++ {

		writer.WriteString(wrtContent)

	}

	//writer 是帶緩存的，其實是先寫到混存裡面，要Flush才真正寫入文件
	writer.Flush()
	fmt.Println("cover succeed")

}

//在原來的內容追加 haha, catch you
func appendAndWrite() {
	fileName := "C:/Users/USER/Desktop/文件檔/createAndWrite.txt"

	//開啟文件的模式可用 | 混用    0666 =>windows 沒用
	//追加內容在文末
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("err occured when open file.. msg :%v\n", err)
		return
	}
	//關閉檔案
	defer file.Close()

	wrtContent := "fucking china\r\n" //如果只寫\n 記事本認不出來
	//寫入時是帶緩存的
	writer := bufio.NewWriter(file) //func NewWriter(w io.Writer) *Writer
	for i := 0; i < 5; i++ {

		writer.WriteString(wrtContent)

	}

	//writer 是帶緩存的，其實是先寫到混存裡面，要Flush才真正寫入文件
	writer.Flush()
	fmt.Println("append succeed")

}

//打開一個存在的文件，將內容顯示在console，追加5句taiwan no.1
func readAndAppend() {
	fileName := "C:/Users/USER/Desktop/文件檔/createAndWrite.txt"

	//開啟文件的模式可用 | 混用    0666 =>windows 沒用
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("err occured when open file.. msg :%v\n", err)
		return
	}
	//關閉檔案
	defer file.Close()

	//讀取文件顯示在console
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			//如果獨到文件末尾
			break
		}
		fmt.Print(str)

	}

	wrtContent := "taiwan no.1\r\n" //如果只寫\n 記事本認不出來
	//寫入時是帶緩存的
	writer := bufio.NewWriter(file) //func NewWriter(w io.Writer) *Writer
	for i := 0; i < 5; i++ {

		writer.WriteString(wrtContent)

	}

	//writer 是帶緩存的，其實是先寫到混存裡面，要Flush才真正寫入文件
	writer.Flush()
	fmt.Println("read and append succeed")

}

func writeToOther() {

	fileName := "C:/Users/USER/Desktop/文件檔/createAndWrite.txt"
	//寫入路徑
	wrtTOpath := "C:/Users/USER/Desktop/文件檔/test2.txt"

	byteData, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("err occuer when read msg:", err)
		return
	}
	err = ioutil.WriteFile(wrtTOpath, byteData, 0666)
	if err != nil {
		fmt.Println("err occuer when write msg:", err)
		return
	}
	fmt.Println("write succeed")
}

//判斷文件是否存在
func pathExist(filePath string) (bool, error) {

	_, err := os.Stat(filePath)
	//1. 如果err 是 nil  => 文件存在
	if err == nil {
		return true, nil
	}
	//2. 如果err 用 os.IsNotExist()==true，文件不存在
	if os.IsNotExist(err) {
		return false, nil
	}
	//3. 其他，不確定是否存在
	return false, nil
}

// 創立文件，並寫入

//os 包 的  func OpenFile(name string, flag int, perm FileMode) (file *File, err error)

//name  :  文件名
//flag  :  模式  (可混用)  p.s. O_TRUNC 要小心使用，會清空檔案

//Constants
//flag  種類
// const (
//     O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
//     O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
//     O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
//     O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
//     O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
//     O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
//     O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
//     O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
// )

//FileMode :  只適用linux

//判斷文件是否存在  os 包  func Stat(name string) (fi FileInfo, err error)
//1. 如果err 是 nil  => 文件存在
//2. 如果err 用 os.IsNotExist()==true，文件不存在
//3. 其他，不確定是否存在
