package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum =", cpuNum)

	//可自己設置使用多少CPU
	//版本: 1.8 以前要設置一下 ，1.8以後預設讓程序運行在多個核上
	runtime.GOMAXPROCS(cpuNum - 1) //這個1.8後可不用寫
	fmt.Println("ok")
}

//go 設置運行CPU個數

// runtime 包 func NumCPU() int  => 返回本地机器的逻辑CPU个数
//https://www.itread01.com/content/1545792242.html
//邏輯CPU: 物理cpu個數×cpu核數×2  =>用虛擬技術讓CPU發揮兩倍效能

//runtime 包 func GOMAXPROCS(n int) int =>设置可同时执行的最大CPU数

//版本: 1.8 以前要設置一下 1.8以後預設讓程序運行在多個核上
