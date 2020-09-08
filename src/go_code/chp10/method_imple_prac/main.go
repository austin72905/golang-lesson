package main

import "fmt"

func main() {
	//1. 寫一個方法打印10*8的矩形
	var rec = new(rectangle)
	rec.printRec()
	//2. 寫一個可以計算面積的方法
	recArea := rec.getArea(3.0, 6.5)
	fmt.Println("area = ", recArea)
	//3. 判斷奇偶的方法
	rec.getSinOrD(31)
}

//1.

type rectangle struct {
}

func (rec rectangle) printRec() {

	for a := 0; a < 10; a++ {
		for b := 0; b < 8; b++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//2.
func (rec rectangle) getArea(len float64, width float64) float64 {
	return len * width
}

//3.
func (rec rectangle) getSinOrD(num int) {
	if num%2 == 0 {
		fmt.Println("偶")
	} else {
		fmt.Println("奇")
	}
}

//方法 題目練習
