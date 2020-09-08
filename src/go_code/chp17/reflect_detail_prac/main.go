package main

import (
	"fmt"
	"reflect"
)

func main() {

	num := 10
	//裡面是值拷貝，所以要傳地址才能改變num的值
	reflectChgVal(&num)
	fmt.Println(num)
}

func reflectChgVal(val interface{}) {

	//得到reflect.Value
	rVal := reflect.ValueOf(val)

	fmt.Printf("rVal Kind=%v\n", rVal.Kind()) //ptr
	//改變他的值
	//func (v Value) SetInt(x int64) 一定要先用Elem()拿到指針指向的數才能調用
	rVal.Elem().SetInt(20)
}

//通過反射修改 num int 的值

//func (v Value) Elem() Value
//Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装
