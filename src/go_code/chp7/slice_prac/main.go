package main

import "fmt"

func main() {

	var intArr [5]int = [...]int{1, 22, 33, 6, 9}
	//定義一個切片
	slice := intArr[1:3] //從index =1 到 index = 3 的前一個
	fmt.Println("intArr = ", intArr)
	fmt.Println("slice元素有 = ", slice)
	fmt.Println("slice元素數量 = ", len(slice))
	fmt.Println("slice的容量 = ", cap(slice)) //切片的容量可以動態變化

	//在內存的形式(類似struct)
	//1. 分成3個部分  1)引用第一個元素的地址
	//這兩個地址應該相等
	fmt.Println("intArr[1]地址 = ", &intArr[1])
	fmt.Println("slice[0]地址 = ", &slice[0])

	//slice 是引用類型，所以一旦修改值，原本引用的陣列也會被修改
	slice[1] = 34
	fmt.Println("intArr = ", intArr)
	fmt.Println("slice元素有 = ", slice)

	//2. make([]類型,長度,容量) => 默認值 0
	var slice1 []float64 = make([]float64, 5)
	fmt.Println(slice1)
	slice1[1] = 10
	slice1[3] = 20
	fmt.Println(slice1)
	fmt.Println("slice1的size = ", len(slice1)) //5
	fmt.Println("slice1的cap =", cap(slice1))   //如果容量沒指定就會跟len一樣

	//3. 定義一個切片，指定具體陣列，類似make
	var slice2 []string = []string{"tom", "jack", "mary"}
	fmt.Println(slice2)
	fmt.Println("slice2的size = ", len(slice2))
	fmt.Println("slice2的cap =", cap(slice2))

	//遍歷  for
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v\n", i, slice[i])
	}

	//for-range
	for idx, val := range slice {
		fmt.Printf("idx = %v , val = %v\n", idx, val)
	}

	//切片使用1. 可以簡寫
	slice3 := intArr[:]
	fmt.Println("slice3 =", slice3)

	//切片可以繼續切片 (跟silce3一樣指向同樣的數據空間)，只是引用的具體位置不同
	slice4 := slice3[0:3]
	fmt.Println("slice4 =", slice4)
	slice4[1] = 30
	//因為指向相同數據空間，一旦值改變其他也會跟著改變
	fmt.Println("intArr = ", intArr)
	fmt.Println("slice3 =", slice3)
	fmt.Println("slice4 =", slice4)

	//可以對切片動態增加 append  (內存會新建新陣列，然後更換引用)
	slice3 = append(slice3, 11, 55, 33)

	fmt.Println("slice3 =", slice3)

	//可以將其他切片合成
	slice3 = append(slice3, slice3...)

	fmt.Println("slice3 =", slice3)

	//拷貝切片
	copy(slice3, slice) //複製切片的值，對應到相對的index

	fmt.Println("slice3 =", slice3)

	//練習
	//將非波納契數 放到切片中
	//非波納契數: 1,1,2,3,5,8,13...
	//=> 1. fn(n)=fn(n-1)+fn(n-2)
	//   2. fn(0)=1
	//   3. fn(0)=1
	fbnSlice := fbn(10)
	fmt.Println(fbnSlice)

}

//返回一個uint64 的切片
func fbn(n int) []uint64 {
	fbnSlice := make([]uint64, n)
	//固定前兩個都是1
	fbnSlice[0] = 1
	fbnSlice[1] = 1

	for i := 2; i < n; i++ {
		//fn(n)=fn(n-1)+fn(n-2)
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

//切片(slice)

//切片是引用類型

//數量不確定的陣列(動態變化)

//定義一個切片

//在內存的形式(類似struct)
//1. 分成3個部分  1)引用第一個元素的地址 2)切片長度 3)切片容量(可選)

// type silce struct{
// 		ptr *[2]int
// 		len int
// 		cap int
// }

//slice 是引用類型，所以一旦修改值，原本引用的陣列也會被修改

//切片的使用
//1. 定義一個切片，去引用陣列
//2. make([]類型,長度,容量) => 默認值 數字 : 0 ,str : 0  ,bool : false
//3. 定義一個切片，指定具體陣列，類似make

//使用2. make 創建切片 內存形式: 會生成一組不可見的陣列，而切片的地址指向他

//遍歷  for 、 for-range

//切片使用1. 可以簡寫   slice:=intArr[f:e]
//當 f=0 、e=len(intArr) 可簡寫，如果要引用整個陣列可寫 slice:=intArr[:]

//切片可以繼續切片

//可以對切片動態增加 append     func append(slice []Type, elems ...Type) []Type 若它有足够的容量，其目标就会重新切片以容纳新的元素。否则，就会分配一个新的基本数组

//可以將其他切片合成

//拷貝切片 copy()  (都是切片)  copy(要貼上的, 要copy)    要copy長度>要貼上的 沒關係，只會修改到對應的index

//sting 和 slice 關係
