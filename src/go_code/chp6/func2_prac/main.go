package main

import "fmt"

//go 裡面  函數也是一種數據類型，也可以當作參數
func main() {
	a := getSum
	//這樣賦值之後 可以調用 a()
	fmt.Printf("a的類型: %T,getSum的類型: %T\n", a, getSum)
	b := a(1, 4)
	fmt.Println(b)
	//函數當成參數
	c := fun(getSum, 5, 6)
	fmt.Println(c)
	//可以自訂一數據類型
	type myint int
	var d myint = 40
	fmt.Printf("%T\n", d) //main.myint
	fmt.Println(d)
	//案例2
	e := fun2(getSum, 5, 6)
	fmt.Println(e)

	//支持對涵式返回值命名
	f, g := getSum2(4, 3)
	fmt.Println(f, g)

	//支持可變參數
	h := sum1(1, 3, 5)
	fmt.Println(h)

}

//案例2
type myFun func(int, int) int //myFun 變成可以代替一個有兩個int參數、回傳 int 的涵式
//如果是自定義 func 數據類型 ，放到main外面才不會報錯

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

//也可以當作參數
func fun(afunc func(int, int) int, num1 int, num2 int) int {
	return afunc(num1, num2)

}

//用自定義的涵式數據 取代func(int, int) int
func fun2(afunc myFun, num1 int, num2 int) int {
	return afunc(num1, num2)

}

//支持對涵式返回值命名
func getSum2(n1 int, n2 int) (sum int, sub int) {
	//返回值已經定義變數了 所以不要再 :=

	sum = n1 + n2
	sub = n1 - n2
	return
}

//支持可變參數
//...  為固定寫法 args 可自己取名，如果可以不傳參數 n1 int 可以省略
//args ...int 一定要放在參數列表最後
//args 本身是切片
func sum1(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

//go 裡面
//函數也是一種數據類型，也可以當作參數
//可以自訂一數據類型
//type myint int  =>myint 可以當作int 用 ，但go 還是認為他不是一樣的數據類型
//支持對涵式返回值命名
//支持可變參數
