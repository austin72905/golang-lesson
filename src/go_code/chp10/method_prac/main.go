package main

import "fmt"

func main() {
	var p Person
	p.Name = "tom"
	//調用方法
	//test()只能通過Person類型的變數來調用
	p.test()
	//方法裡是值拷貝，所以不會影響到外面的p.Name，除非用指針
	fmt.Println("name =", p.Name)

	//練習
	//1. 聲明一個結構體Cycle，屬性為radius
	//2. 聲明一個方法area和Circle綁定，可以返回面積
	cirlce := Circle{3.1}
	area := cirlce.area()
	fmt.Println("area = ", area)

	//如果type 是指針 標準的調用方法
	area2 := (&cirlce).area2() //等於這樣寫(編譯器幫你優化了) cirlce.area2()
	fmt.Println("area2 = ", area2)

	//呼叫一個自訂一類型的方法
	var inter myInter = 10
	inter.printPlus()
	fmt.Println("自定義方法值 =", inter) // 11

	//自定義類型實現了String()，之後fmt.Println會默認調用這個變數的String()進行輸出
	myStu := Student{
		Name: "austin",
		Age:  16,
	}

	fmt.Println(&myStu) //名字= austin ,年紀 = 16

	//.  方法的參數 是接收 "值類型" ，可以用指針類型的變數調用方法，反之亦然
	(&inter).printInt() //printInt()這個方法是接收(inter myInter) ，但可以用(&inter)呼叫，因為編譯器自動優化了，本質仍是值拷貝
	//最終決定權是看你 方法 定義傳進去的是值類型還是引用類型 ， (inter myInter)  或  (inter *myInter)

	//也可以在宣告時把一個strcut的指針返回給指針變數
	per2 := &Person{"austin"}
	//會打印出地址
	fmt.Println(per2) //&{austin}

}

//Person struct
type Person struct {
	Name string
}

//聲明一個綁定Person的方法
func (p Person) test() {
	p.Name = "jack" //是值拷貝，所以不會影響到外面的p.Name，除非用指針
	fmt.Println("test() name =", p.Name)
}

//Circle struct
type Circle struct {
	radius float64 //半徑
}

func (c Circle) area() float64 {
	return c.radius * c.radius * 3.14
}

//通常為了效率會和結構體的指針綁定
func (c *Circle) area2() float64 {
	//標準應該要(*c).radius 才取的到值，但go底層幫你轉好了
	return c.radius * c.radius * 3.14
}

//golang中方法，不一定要綁定struct ，只要是自定義的數據類型，都能有方法
//自定義一個int
type myInter int

func (inter *myInter) printPlus() {
	*inter = *inter + 1
}

//Student struct
type Student struct {
	Name string
	Age  int
}

func (stu *Student) String() string {
	defStr := fmt.Sprintf("名字= %v ,年紀 = %v", stu.Name, stu.Age)
	return defStr
}

//.  方法的參數 是接收 "值類型" ，可以用指針類型的變數調用方法，反之亦然
func (inter myInter) printInt() {
	fmt.Println(inter)
}

//方法

//1. golang中方法，不一定要綁定struct ，只要是自定義的數據類型，都能有方法

//2. 聲明方法

//3. 方法傳遞參數的機制: 他會同時複製一份類型(p Person) 變數進來(如果變數類型是值類型=>值拷貝、如果是引用類型=>地址拷貝)，與涵式不同

//4. 通常為了效率會和結構體的指針綁定

//5. 如果一個自定義類型實現了String()，之後fmt.Println會默認調用這個變數的String()進行輸出(輸出日誌常用)

//6. 方法與函數的差異
//    a. 一般的函數，如果其參數 是接收 "值類型"，呼叫的時候就只能給他值類型 ，如果是"引用"，也只能給引用
//    b.  方法的參數 是接收 "值類型" ，可以用指針類型的變數調用方法，反之亦然
//    c. 調用方法不一樣，方法要先定義一個變數 then  變數.方法

//7. 也可以在宣告時把一個strcut的指針返回給指針變數
