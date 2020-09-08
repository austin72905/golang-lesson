package main

import "fmt"

func main() {
	//test(4)
	//test2(4)
	//res := fbn(6)
	//fmt.Println("res=", res)
	ans := peach(1)
	fmt.Println("ans=", ans)
}

//遞迴 函數調用自己本身

func test(n int) {
	if n > 2 {
		n-- //如果不加這條件會變成無限迴圈
		test(n)
	}
	fmt.Println(n)
}

//進來就變 3
//=>test(3) => 進來變2
//=> test(2) =>print 2 (test(2)印的)
//=> print 2 (test(3)印的)=>
//print 3 (test(4)印的)

func test2(n int) {
	if n > 2 {
		n--
		test(n)
	} else {
		fmt.Println(n)
	}

}

//test2(3)=>test(2)=>print 2 (test2(2)印的)

//練習  非波納契數1,1,2,3,5,8,13
//思路:
//1. 當 n=1 或 n=2  ; 返回1
//2. 當 n>=2 ; 返回前兩個數的和 f(n-1)+f(n-2)

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	//go會不給你用else，因為這邊用else是多餘的
	return fbn(n-1) + fbn(n-2)

}

//練習
// 有一堆 桃子 ，猴子第一天 吃了其中的一半多一顆， 之後的每天 都吃其中的一半多一顆
//到了第 10天 ，還沒吃，發現剩下1顆桃子， 請問最初有幾顆?

//找規律!!!!!!!!!找規律!!!!!!!找規律!!!!!!!!!!!找規律!!!!!!!!!!!!!!!
//程式語言就是找規律
//第10 天:  1 個桃子
//第9天  :  (第10天的桃子數量+1)*2
//第8天    (第9天的桃子數量+1)*2
//規律  => (前一天的桃子數+1)*2
//peach(n)=(peach(n+1)+1)*2

// n => 第幾天  返回 剩幾個桃子
func peach(n int) int {
	if n == 10 {
		return 1
	}
	return (peach(n+1) + 1) * 2
}
