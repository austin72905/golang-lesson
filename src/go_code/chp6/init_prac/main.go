package main

import "fmt"

//init 涵式可以在每個源文件中存在一個，他會在main執行前 ，先執行init
//可以做一些初始化工作，如在要引入的包 init() 初始化要引入變數的值
//如果文件裡面有全局變數、init、main  執行順序:  全局變數 => init => main
//如果 main.go 跟 utils.go 都有全局變數跟init ，會先執行utils.go 的
func init() {
	fmt.Println("init()...")
}
func main() {
	fmt.Println("main()..")
}
