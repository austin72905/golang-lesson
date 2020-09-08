package main

import "fmt"

func main() {
	computer := computer{}
	phone := phone{}
	camera := camera{}

	//能夠傳進去working的struct參數一定是要實作介面所有方法的，否則報錯
	computer.working(phone)
	//cannot use phone (type phone) as type usb in argument to computer.working:
	//phone does not implement usb (missing stop method)
	computer.working(camera)

	//多載陣列
	var usbArr [3]usb
	//這樣寫會報錯 phone is not a type
	// usbArr[0] = phone{}
	// usbArr[1] = phone{}
	// usbArr[2] = camera{}
	//不知道是不是因為前面實例化過了
	usbArr[0] = gopro{"gopro"}
	usbArr[1] = gopro{"gopro7"}
	usbArr[2] = recorder{"deina"}
	fmt.Println(usbArr)
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
	usb.stop()
}

type gopro struct {
	name string
}

func (g gopro) start() {
	fmt.Println("gopro loading...")
}

func (g gopro) stop() {
	fmt.Println("gopro sleeping...")
}

type recorder struct {
	name string
}

func (r recorder) start() {
	fmt.Println("recorder loading...")
}

func (r recorder) stop() {
	fmt.Println("recorder sleeping...")
}

//接口(介面)

//定義一個規範，讓其他結構體來實作他

//go的介面不能有變數，介面裡面可以定義多個方法(都沒有方法體)

//go的介面不需要顯式實現(不用implement、expend...)，只要有一個變數，含有介面的"所有"方法，這個變數就實現了這個介面

//多載陣列: 透過介面的方式
