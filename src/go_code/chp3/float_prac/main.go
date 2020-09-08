package main

import "fmt"

func main() {
	var price float32 = 3.1415
	fmt.Println(price)
	//尾數部分可能丟失，精度損失
	//儘量使用float64
	var a float32 = -123.0000901 //後面的01會丟失
	var b float64 = -123.0000901
	fmt.Println(a, b)
	var c = 3.41
	fmt.Printf("c的類型   %T\n", c)

	//十進制 5.12
	d := 5.12
	e := .512
	fmt.Println(d, e)
	//科學計數法
	f := 1.155e2  // 1.155 * 10^2
	g := 1.155e-2 // 1.155 / 10^2
	fmt.Println(f, g)

}

//float32   占用: 4字節
//float64   8字節 ->類似其他語言的double

//1. 都是有符號
//2. 尾數部分可能丟失，精度損失
//3. 儲存 符號位+指數位+尾數位  (計個大概)
//4. 預設float64
//5. 儘量使用float64
