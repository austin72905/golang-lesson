package main

import (
	"flag"
	"fmt"
)

func main() {
	//new 一個strcut
	var userInfo userInfo

	//&userInfo.user 接收  命令行 -u 後面的參數值
	//"u" , 就是-u 指定參數
	//"" , 默認值
	//"用戶名，默認為空" , 說明
	flag.StringVar(&userInfo.user, "u", "", "用戶名，默認為空")
	flag.StringVar(&userInfo.password, "pwd", "", "密碼，默認為空")
	flag.StringVar(&userInfo.host, "h", "localhost", "主機名，默認localhost")
	flag.IntVar(&userInfo.port, "port", 8080, "端口，默認8080")

	//必須調用這個方法才能夠操作
	flag.Parse()

	//輸出結果
	fmt.Printf("user=%v\n", userInfo.user)
	fmt.Printf("password=%v\n", userInfo.password)
	fmt.Printf("host=%v\n", userInfo.host)
	fmt.Printf("port=%v\n", userInfo.port)
}

type userInfo struct {
	user     string
	password string
	host     string
	port     int
}

//flag 包  func IntVar(p *int, name string, value int, usage string)

//如果打錯會顯示詳細訊息超帥~~

// flag provided but not defined: -password
// Usage of loginTest.exe:
//   -h string
//         主機名，默認localhost (default "localhost")
//   -port int
//         端口，默認8080 (default 8080)
//   -pwd string
//         密碼，默認為空
//   -u string
//         用戶名，默認為空
