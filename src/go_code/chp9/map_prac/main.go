package main

import "fmt"

func main() {

	//宣告方法
	var frtMap map[string]string
	//直接賦值會報錯，要make 初始化才能使用
	frtMap = make(map[string]string, 10) //可以放 10 pairs

	frtMap["first"] = "one"
	fmt.Println(frtMap)

	//宣告方法2   : 直接make
	secMap := make(map[string]string) //size 可忽略
	secMap["1"] = "a"
	secMap["2"] = "b"
	fmt.Println(secMap)

	//宣告方法3   :  直接賦值
	thdMap := map[string]string{
		"aa": "aa",
		"bb": "bb",
	}
	fmt.Println(thdMap)

	//練習
	//紀錄3個學生訊息，每個學生有name 、gender
	//valType =>可以用map
	forthMap := make(map[string]map[string]string)
	//這樣寫也可以
	// forthMap["stud1"] = make(map[string]string)
	// forthMap["stud1"]["name"]="tom"
	forthMap["stud1"] = map[string]string{
		"name":   "austin",
		"gender": "boy",
	}

	forthMap["stud2"] = map[string]string{
		"name":   "mimi",
		"gender": "girl",
	}
	fmt.Println(forthMap)
	fmt.Println(forthMap["stud1"]["name"])

	//map 的 crud
	//新增、修改  :  如果key 不存在就增加，存在就修改
	//刪除   :  如果key 不存在就不操作(不會報錯)，存在就刪除
	thdMap["cc"] = "cc"
	fmt.Println(thdMap)

	delete(thdMap, "cc")
	fmt.Println(thdMap)

	//刪除全部的key 1.歷遍，逐一刪除 2. make 一個新的 讓gc回收

	//查詢 key 是否存在
	//findResult  bool  key存在 : true ,key不存在 : false
	//val  查詢key的value
	val, findResult := thdMap["cc"]
	if findResult {
		fmt.Printf("key存在 , 其值為%v\n", val)
	} else {
		fmt.Printf("key不存在 , 其值為%v\n", val)
	}

	//map的遍歷 : 只能用for-range

	for key, value := range thdMap {
		fmt.Printf("key = %v , value = %v\n", key, value)
	}
	//較複雜的結構
	for key, value := range forthMap {
		fmt.Printf("%v\n", key)
		//forthMap的value是map
		for keyIsd, valIsd := range value {
			fmt.Printf("  key: %v , val : %v\n", keyIsd, valIsd)
		}
	}

}

//map

//類似C#的 Dic    key不能重複、沒有順序(每次遍歷會不一樣)

//宣告方法 var  變數  map[keyType]valType

//除了silce 、 map 、 func  其他都可以當keyType

//宣告不會分配內存，直接賦值會報錯，要make 初始化才能使用，跟陣列不同

//map 的 crud

//map的遍歷 : 只能用for-range

//map的長度 len  可察看有幾pairs

//map 切片
