package main

import (
	"fmt"
	"strconv"
)

//如果沒使用到一個包 _ "fmt" =>暫時忽略
func main() {
	//數據轉換練習
	//T(v)
	var a int = 1
	var b float32 = float32(a)
	fmt.Printf("%v %v\n", a, b)
	//這個b =>1.000000
	fmt.Printf("%v %f\n", a, b)
	//變換的變數的"值"，變數本身沒有發生變化
	fmt.Printf("a type is %T\n", a)

	//溢出處理
	//轉換的數據要是該數據範圍內的
	var c int64 = 999999
	var d int8 = int8(c) //63
	fmt.Println(d)

	//練習
	var e int32 = 12
	var f int64
	// f = e + 12  這樣寫會報錯，因為後面的 e + 12 =>int32， 與f  (int64) 不能直接轉換
	f = int64(e) + 12
	fmt.Println(f)

	//其他類型轉string
	var g int = 9
	var h float64 = 3.145
	var i bool = true
	var j byte = 'j'
	var str string

	//int =>string
	str = fmt.Sprintf("%d", g)
	fmt.Printf("k type %T k=%q\n", str, str) //%q 有雙引號
	//float64 =>string
	str = fmt.Sprintf("%f", h)
	fmt.Printf("k type %T k=%q\n", str, str)
	//bool =>string
	str = fmt.Sprintf("%t", i)
	fmt.Printf("k type %T k=%q\n", str, str)
	//byte => string
	str = fmt.Sprintf("%c", j)
	fmt.Printf("k type %T k=%q\n", str, str)

	fmt.Println("--------------第二種方法---------------")
	//第二種方式
	str = strconv.FormatInt(int64(g), 10) //func FormatInt(i int64, base int) string  =>後面代表轉成幾進制
	fmt.Printf("k type %T k=%q\n", str, str)

	//func ParseFloat(s string, bitSize int) (f float64, err error)
	//'f' : 格式 10 : 表示保留10位  64 : float64
	str = strconv.FormatFloat(h, 'f', 10, 64)
	fmt.Printf("k type %T k=%q\n", str, str)

	str = strconv.FormatBool(i)
	fmt.Printf("k type %T k=%q\n", str, str)

	fmt.Println("--------------第三種方法---------------")
	//整數轉字串
	str = strconv.Itoa(g)
	fmt.Printf("k type %T k=%q\n", str, str)

	fmt.Println("--------------str 轉 其他類型---------------")
	//string 轉 其他類型
	var wstr string = "true"
	var boln bool
	//func ParseBool(str string) (value bool, err error) 會返回兩個值
	// _ 忽略 err
	//返回都是64 所以都要用64接
	boln, _ = strconv.ParseBool(wstr)
	fmt.Printf("boln type %T boln=%v\n", boln, boln)
	//func ParseInt(s string, base int, bitSize int) (i int64, err error)
	//bitSize int => 0 int ,8 int8 以此類推
	var wstr2 string = "123456"
	var strInt int64
	strInt, _ = strconv.ParseInt(wstr2, 10, 64)
	fmt.Printf("strInt type %T strInt=%v\n", strInt, strInt)
	//func ParseFloat(s string, bitSize int) (f float64, err error)
	var wstr3 string = "123.56"
	var strFloat float64
	strFloat, _ = strconv.ParseFloat(wstr3, 64)
	fmt.Printf("strFloat type %T strFloat=%v\n", strFloat, strFloat)
	//如果不能轉ex: hello =>會被轉成 0
	var wstr4 string = "hello"
	var strInt0 int64
	strInt0, _ = strconv.ParseInt(wstr4, 10, 64)
	fmt.Printf("strInt0 type %T strInt0=%v\n", strInt0, strInt0)
}

//數據轉換
//1. 範圍大到小、小到大都可
//2. 變換的變數的"值"，變數本身沒有發生變化
//3. int64轉int8 編譯不會抱錯，但轉換結果按溢出處理，結果會和預期不同

//其他類型轉string
//1. fmt.Sprintf("%參數",表達示) =>格式化字符串
//2. 使用strconv 的包

//string 轉 其他類型
//1. 如果不能轉 int ex: hello =>會被轉成 0
//2. 如果不能轉 bool => false
