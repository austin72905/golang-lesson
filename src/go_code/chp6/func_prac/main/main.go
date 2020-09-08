package main

import (
	"fmt"
	"go_code/chp6/func_prac/tool" //$GOPATH到src都可以省略
	//toolbag  "go_code/chp6/func_prac/tool"
)

func main() {
	//練習 寫一個計算機
	// result := tool.Cal()         //也可以在路徑前面加個別名toolbag.Cal()
	// fmt.Printf("%.2f\n", result) //印出到小數第二位

	//如果只想要一個值可用 _ 忽略
	sum, sub := tool.GetSum(5, 4)
	fmt.Println(sum, sub)

	//如果希望函數內的變數，能修改函數外的變數，可以傳入變數的地址&
	num := 20
	test3(&num)
	fmt.Println("main() num=", num)

}

//他是指針類型的參數
//n 是指到 num所在的地址
//*n 指到所在地址(num)裡面的值
func test3(n *int) {
	*n = *n + 10 //直接修改了num的值了，就算函數調用完會被回收，但值已被修改
	fmt.Println("test3() n=", *n)
}

//函數
//1. return 可以返回值，也代表終止涵式
//2. 如果希望函數內的變數，能修改函數外的變數，可以傳入變數的地址&
//3. 不支持多型  ，go 是使用空接口的方式
//4. 值傳遞與引用傳遞(傳地址、效率高)
//5. 值傳遞: int、float、bool、string、數組、struct
//6. 引用傳遞: 指針、silce、map、chan、interface
//7. 變數:=值 (賦值語句不能再函數外執行)

//包  注意目錄
//每個文件都屬於一個包
//區分相同名稱函數
//文件夾跟(包名)不一定要一樣 ，但引用涵式要記得死用包名
//大寫涵式才可被其他文件引用
//如果要編譯成一個可執行文件，就需要該包聲明為main

// 編譯go 文件
// 到 $GOPATH 所在目錄 (C:\Users\USER\GoPro)
// 指令 go build  go_code\chp6\func_prac\main    ($GOPATH/src之後的檔案路徑)
//    go build -o bin\mypro.exe go_code\chp6\func_prac\main
//    go build -o {當前目錄的檔案}\{自訂義檔名}.exe {$GOPATH/src之後的檔案路徑}

// 庫文件
// pkg 裡面有個a檔案  是二進制檔案
// 給人庫文件，別人也能引用，可以不用看原始碼

//涵式底層運作
//棧區 =>空出一個 main 棧區
//調用Cal =>空出一個 Cal 棧區 =>調用完畢就被編譯器自動回收
