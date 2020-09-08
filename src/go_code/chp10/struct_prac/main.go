package main

import "fmt"

func main() {
	//建立一個Cat 變數
	var cat1 Cat //會有默認值
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	fmt.Println("貓貓訊息如下")
	fmt.Println("name =", cat1.Name)
	fmt.Println("age = ", cat1.Age)
	fmt.Println("color =", cat1.Color)
	fmt.Printf("cat1 地址 = %p\n", &cat1)

	//宣告變數即有默認值
	var p1 Person
	fmt.Println(p1) //{ 0 [0 0] <nil> [] map[]}
	//sliece、map =>nil(還沒有分配空間)需要先make才能用
	p1.slice = make([]int, 3)
	p1.slice[0] = 10
	p1.map1 = make(map[string]string)
	p1.map1["fru1"] = "apple"
	fmt.Println(p1) //{ 0 [0 0] <nil> [10 0 0] map[fru1:apple]}

	//struct 間值不會互相影響(struct 值類型)
	fmt.Println("struct 間值不會互相影響(struct 值類型)")
	p2 := p1
	fmt.Println(p1)
	fmt.Println(p2)
	p2.Name = "dick"
	fmt.Println(p1)
	fmt.Println(p2)
	//要修改用指向地址

	//strcut 宣告方式
	//方法一
	fmt.Println("-----方法一-----")
	p3 := Person{}
	p3.Name = "tom"
	fmt.Println(p3)
	//方法二
	fmt.Println("-----方法二-----")
	cat2 := Cat{"org", 2, "orange"}
	fmt.Println(cat2)
	//方法三
	//返回的是一個指針
	fmt.Println("-----方法三-----")
	var cat3 *Cat = new(Cat)
	//賦值方式  go 底層會進行優化  cat3.Name = "ha" 也可以賦值
	//記得(*cat3).Name   *cat3.Name=>會報錯 因為. 優先執行於*
	(*cat3).Name = "ha"
	cat3.Color = "HAB"

	fmt.Println(cat3) //&{ha 0 HAB}      cat3 指向的地址
	fmt.Printf("cat3 指向的地址 =%p\n", cat3)

	fmt.Println(&(*cat3)) //&{ha 0 HAB}      cat3 指向的值所在的地址
	fmt.Printf("cat3 指向的值所在的地址 =%p\n", &(*cat3))

	fmt.Println(&cat3) //0xc0000ca020  cat3這個指針本身的地址

	fmt.Println((*cat3)) //{ha 0 HAB}       cat3 指向的值

	//方法四
	fmt.Println("-----方法四-----")
	var cat4 *Cat = &Cat{} //也可以在{}裡面給值
	(*cat4).Name = "bo"
	cat4.Color = "bokesi"
	fmt.Println(cat4)
	fmt.Println(&cat4)
	fmt.Println(&(*cat4))
	fmt.Println((*cat4))

}

//Cat struct
//定義一個 CAT 類
//大寫可被其他包引用
type Cat struct {
	Name  string
	Age   int
	Color string
}

//Person struct
type Person struct {
	Name  string
	Age   int
	Score [2]float64
	ptr   *int              //指針
	slice []int             // 切片
	map1  map[string]string //map
}

//物件導向(OOP)(面向對像)  對象=>在go裡面struct   、其他語言 class

//方便管理數據(都放在struct 的屬性)

//go 沒有class ，用struct 代替，支持OOP特性，但不是純粹的OOP (有全局變數)，其他語言遍數都在class裡面

//go的OOP很簡潔 沒有繼承(跟其他語言的繼承不一樣)、多載、this、建構子

//還是有OOP的繼承、封裝、多態，只是實現方式不一樣

//通過接口(interface)關聯 ， 耦合性低

//strcut  1. 屬性  2. 行為(方法)，通過strcut 可以創建很多實例(變數)

//strcut (值類型)

// 1. 本身是自定義的數據類型

// 2. 宣告後的strcuc變數 是一個實例

// 3. 內存
//    a. 宣告變數後，變數會指向該struct (struct裡面有幾個屬性，內存就會有幾個字段(屬性)(field))
//    b. 賦予屬性值的時候，變數就會找到指定的字段，給值 ex: cat1.Name = "小白"
//    c. struct 裡面如果有slice、map 都是存地址，再指向其他空間
//    d. 裡面的屬性在內存是連續分布的

// 4. 宣告變數即有默認值   str =""、bool =false(有分配空間)  ptr(用 (基本數據類型)new 才能使用) ，sliece、map =>nil(還沒有分配空間)需要先make才能用

// 5. struct 間值不會互相影響(struct 值類型)

// 6. strcut 宣告方式
