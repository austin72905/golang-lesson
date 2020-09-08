package main

import "fmt"

//同一個資料夾，只能有一個main
func main() {
	fmt.Println("austin\tlin")                                 // \t =>空格
	fmt.Println("chg\nline")                                   // \n =>換行
	fmt.Println("C:\\Users\\USER\\GoProj\\src\\go_code\\chp2") // \\ =>如果要印出的文字有\ 可用到
	fmt.Println("austin says \"I LOVE YOU\"")                  // 要印出""也行
	fmt.Println("I AM your mom\rYOUR DAD")                     // \r 迴車 \r後面會從頭把原本的文字覆蓋掉
	//練習
	fmt.Println("姓名\t年齡\t籍貫\t住址\njohn\t12\t河北\t北京")

}
