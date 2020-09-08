package main

import "fmt"

func main() {
	//引用類型，用一個函數接收map ，修改後會修改原來的map
	intMap := map[int]string{
		1: "jan",
		2: "feb",
		3: "mar",
		4: "apr",
		5: "may",
	}
	modify(intMap)
	fmt.Println(intMap)
	//map 的value 可以用 struct

	stu1 := studInfo{"austin", 25}
	stu2 := studInfo{"yoyo", 15}

	stuCls := map[string]studInfo{
		"classA": stu1,
		"classB": stu2,
	}
	fmt.Println(stuCls)
	//取出訊息
	for key, studInfo := range stuCls {
		fmt.Printf("班級 : %v\n", key)
		fmt.Printf("姓名 : %v\n", studInfo.Name)
		fmt.Printf("年齡 : %v\n", studInfo.age)
		fmt.Println()
	}
}
func modify(intMap map[int]string) {
	intMap[5] = "May is snow"
}

//結構體
type studInfo struct {
	Name string
	age  int
}

//map 細節

//引用類型，用一個函數接收map ，修改後會修改原來的map

//map容量到達後，想再增加元素，會自動擴充，不會panic

//map 的value 可以用 struct
