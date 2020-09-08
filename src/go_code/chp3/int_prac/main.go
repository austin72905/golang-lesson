package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int8 = 127
	fmt.Println(a)
	var b uint8 = 255
	fmt.Println(b)
	var c rune = '我' //25105
	fmt.Println(c)
	var d byte = 'a' //97
	fmt.Println(d)

	//不給類型預設int
	var e = 100
	//格式化輸出
	fmt.Printf("e 的類型 %T \n", e) //%T 不加繪輸出怪怪的
	var f int64 = 10
	//%d 輸出一個整數
	//unsafe.Sizeof(f) : unsafe包裡的一個func ，可返回占用的字節數
	fmt.Printf("f 的類型 %T  占用的字節數 %d", e, unsafe.Sizeof(f))
	//開發原則: 儘量用小的
	//很小的整數可用byte 0~255
}

//有符號
//int8   占用1 字節 -128~127
//int16  2 字節
//int32  4 字節
//int64  8 字節

//無符號  : 有符號的第一個字節要放符號，所以一個字節只能存7個
//uint8  同上  0~255
//uint16
//uint32
//uint64

//rune 有符號 等價int32  表是一個unicode碼
//byte 無符號 等價uint8  存單個字符使用他 0~255
