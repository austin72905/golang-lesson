package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a = false
	fmt.Println(a, "佔用空間", unsafe.Sizeof(a))

}

//bool 只佔一個字節
