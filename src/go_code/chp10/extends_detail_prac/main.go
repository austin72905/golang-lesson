package main

import "fmt"

func main() {
	var res resP
	//res.MypayURL 這樣寫會報錯， 訪問有名結構體要明確
	res.respData.MypayURL = "http://google.com"
	fmt.Println(res)

	//struct變數宣告時，可以同時宣告匿名結構體的值
	//userInfo := sale{userInfo{"ausit", "63"}, buyItem{"6000", 1}}
	userInfo := sale{
		userInfo{
			name: "ausit",
			age:  "63",
		},
		buyItem{
			price: "6000",
			count: 1,
		}}
	fmt.Println(userInfo)

}

type resP struct {
	code     string
	msg      string
	respData data //有名結構體

}

type data struct {
	MypayURL string //MypayUrl 不能寫這樣真奇怪..
}

//struct變數宣告時，可以同時宣告匿名結構體的值
type sale struct {
	userInfo //匿名結構體
	buyItem
}

type userInfo struct {
	name string
	age  string
}

type buyItem struct {
	price string
	count int
}

//繼承細節

//可以同個struct 繼承 多個 匿名結構體

//struct choose同時繼承 A、B兩個，
//A、B 剛好都有Name這個屬性，而choose 沒有，這時要明確指定結構體名字 ，否則編譯器不知選哪個

//組合: 繼承一個有名的結構體 (類似 C#  class 裡 又放一個class)， 要明確指定結構體名字

//struct變數宣告時，可以同時宣告匿名結構體的值

//結果體屬性可以單純是數據類型

//多重繼承(建議不使用，保持簡潔性)
