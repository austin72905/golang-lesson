package main

import (
	"fmt"
	"os"
)

func main() {
	//練習
	fmt.Println("輸入的參數有", len(os.Args))

	//歷遍os.Args切片 ，獲取命令行的參數
	for index, val := range os.Args {
		fmt.Printf("args[%v]=%v\n", index, val)
	}
}

//命令行參數

//獲取在 exe 檔 後面輸入的參數  (以空格分隔)

// 使用 os 包 裡 Args  方法  他是一個 []string

//powershell預設並不會從目前的位置載入命令    開啟exe要.\{檔名}

//指令: 到exe檔所在文件  檔名.exe 參數 參數二(以空格分)

//缺點: 要按照參數順序輸入

//解決: 使用flag 包  (參數順序可以隨意)
