package main

import "fmt"

func main() {
	//多載陣列
	var usbArr [3]usb

	usbArr[0] = phone{}
	usbArr[1] = phone{}
	usbArr[2] = camera{}

	fmt.Println(usbArr)
	var com computer
	//歷遍usbArr ，把不同的結構體分別放入working
	for _, val := range usbArr {
		com.working(val)
	}
	//練習: 寫一個函數，判斷傳入的參數是什麼類型
	stud1 := student{}
	stud2 := &student{}
	typejudge("austin", 1, 1.6, stud1, stud2, com)
}

//定義一個接口(介面)
//ex: 電腦有USB接口，根據你插入的是IPHONE 還是 CAMERRA 分別 來 實作

type usb interface {
	start()
	stop()
}

//分別定義手機、相機來實作
type phone struct {
}

func (p phone) start() {
	fmt.Println("phone loading...")
}

func (p phone) stop() {
	fmt.Println("phone sleeping...")
}

//phone 獨有的方法
func (p phone) call() {
	fmt.Println("phone calling...")
}

type camera struct {
}

func (c camera) start() {
	fmt.Println("camera loading...")
}
func (c camera) stop() {
	fmt.Println("camera sleeping...")
}

//電腦struct
type computer struct {
}

//定義一個方法，接收介面
//這也算是多載了，依據傳入的參數，決定要實作哪個usb.start()、usb.stop()

func (com computer) working(usb usb) {
	//通過介面來調用方法
	usb.start()

	//如果usb 指向phone 結構體變數，則還需要調用call()
	//使用類型斷言，如果usb能夠轉成phone ，就調用phone裡面的call方法
	if usbPtr, isPhone := usb.(phone); isPhone {
		usbPtr.call()
	}

	usb.stop()
}

func typejudge(itmes ...interface{}) {
	for index, arg := range itmes {
		switch arg.(type) { //type 是關鍵字，固定寫法
		case bool:
			fmt.Printf("參數 %v is a bool 值 = %v\n", index, arg)
		case float64:
			fmt.Printf("參數 %v is a float64 值 = %v\n", index, arg)
		case int, int64:
			fmt.Printf("參數 %v is a int 值 = %v\n", index, arg)
		case string:
			fmt.Printf("參數 %v is a string 值 = %v\n", index, arg)
		case nil:
			fmt.Printf("參數 %v is a nil 值 = %v\n", index, arg)
		case student:
			fmt.Printf("參數 %v is a student 值 = %v\n", index, arg)
		case *student:
			fmt.Printf("參數 %v is a *student 值 = %v\n", index, arg)
		default:
			fmt.Printf("參數 %v is unknown 值 = %v\n", index, arg)
		}
	}
}

type student struct {
}

//類型斷言

//如果working() 傳入的是phone 調用它特有的call()方法

//練習: 寫一個函數，判斷傳入的參數是什麼類型
