package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//裡面的屬性在內存是連續分布的
	r1 := Rect{Point{1, 2}, Point{3, 4}}

	fmt.Printf("r1裡屬性的地址 =%p   %p   %p   %p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y) //int 8個 byte

	//如果屬性是指針，本身地址是連續的，但指向的地址不一定是連續的
	r2 := Rect2{&Point{1, 2}, &Point{3, 4}}

	fmt.Printf("r2裡屬性本身的地址 =%p   %p\n",
		&r2.leftUp, &r2.rightDown) //int 8個 byte
	//但指向的地址不一定是連續的
	fmt.Printf("r2裡屬性指向的地址 =%p   %p\n",
		r2.leftUp, r2.rightDown)

	//struct 轉換
	//裡面的屬性要完全相同
	var a Anum
	var b Bnum
	b = Bnum(a)
	fmt.Println(a, b)

	//struct 可以寫上tag ，該tag可以透過反射獲取，使用場景就是序列化和反序列化
	//將struct序列化成json字串返回時，用戶希望屬性是小寫，可以用到
	monster := Monster{"dog", 3, "bark"}

	//將變數轉成json字串
	//如果把strcut 屬性改小寫，json包抓不到，要在屬性後面加`json:希望顯示的名稱`
	jsonStr, _ := json.Marshal(monster) //返回[]byte

	fmt.Println("monsJson =", string(jsonStr)) //要用string()才能正確顯示

}

//Point struct
type Point struct {
	x int
	y int
}

//Rect struct
type Rect struct {
	leftUp, rightDown Point
}

//Rect2 struct
//如果屬性是指針，本身地址是連續的，但指向的地址不一定是連續的
type Rect2 struct {
	leftUp, rightDown *Point
}

//Anum struct
//struct 轉換
type Anum struct {
	Num int
}

//Bnum struct
type Bnum struct {
	Num int
}

//Monster struct
type Monster struct {
	Name  string `json:"name"` //json是固定的
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

//struct 細節

// 1. 內存
//    a. 宣告變數後，變數會指向該struct (struct裡面有幾個屬性，內存就會有幾個字段(屬性)(field))
//    b. 賦予屬性值的時候，變數就會找到指定的字段，給值 ex: cat1.Name = "小白"
//    c. struct 裡面如果有slice、map 都是存地址，再指向其他空間
//    d. 裡面的屬性在內存是連續分布的
//    e. 如果屬性是指針，本身地址是連續的，但指向的地址不一定是連續的

//2. struct 轉換
//   a. 如果兩個struct 要轉換，裡面的屬性要完全相同

//3. struct進行type重新定義，就會變成新的數據類型，但可以互相強制轉換

//4. struct 可以寫上tag ，該tag可以透過反射獲取，使用場景就是序列化和反序列化 ，將struct序列化成json字串返回時，用戶希望屬性是小寫，可以用到
