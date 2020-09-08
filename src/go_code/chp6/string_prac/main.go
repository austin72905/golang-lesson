package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//取得長度
	str := "hello安"
	fmt.Println("str len= ", len(str)) //8 中文佔3字節

	//處理字串 有中文
	strToRune := []rune(str)
	for i := 0; i < len(strToRune); i++ {
		fmt.Printf("str=%c\n", strToRune[i])
	}

	//字串轉整數
	str2, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("parse err", err)
	} else {
		fmt.Println("parse succeed , val=", str2)
	}

	//整數轉字串
	str3 := strconv.Itoa(135)
	fmt.Printf("str3=%v, TYPE=%T\n", str3, str3)

	//字串轉byte[]
	var bytes = []byte("hello")
	fmt.Printf("bytes=%v\n", bytes) //ascii 碼

	//byte[]轉字串
	str4 := string([]byte{97, 98, 99}) //裡面是大括號
	fmt.Printf("str4=%v\n", str4)

	//10進制轉 2,8,16進制
	str = strconv.FormatInt(123, 2) //2 、8、16
	fmt.Printf("123對應的二進制=%v\n", str)

	//字串是否包含某字串
	isIn := strings.Contains("seafood", "foo")
	fmt.Printf("isIn=%v\n", isIn)

	//計算字串裡有幾個指定的字串
	coutNum := strings.Count("cheese", "e")
	fmt.Printf("coutNum=%v\n", coutNum)

	//不區分大小寫的字串比較
	compStr := strings.EqualFold("ABC", "abc")
	fmt.Printf("compStr=%v\n", compStr)

	//返回某字串在字串第一次出現的索引值，都沒有返回-1
	strIndex := strings.Index("aaaABCaaa", "ABC")
	fmt.Printf("strIndex=%v\n", strIndex)

	//返回某字串在字串最後一次出現的索引值
	strLastIndex := strings.LastIndex("go golang", "go")
	fmt.Printf("strLastIndex=%v\n", strLastIndex)

	//指定字串取代
	strReplace := strings.Replace("go go golang", "go", "美國", -1) //最後參數<0 全部替換
	fmt.Printf("sstrReplace=%v\n", strReplace)

	//分割字串
	strSpt := strings.Split("132,456,789", ",")
	for i := 0; i < len(strSpt); i++ {
		fmt.Println(strSpt[i])
	}

	//大小寫
	strCase := strings.ToUpper("hello")
	fmt.Println(strCase)

	//字串去左右空格
	strTrimS := strings.TrimSpace(" no one hold me ")
	fmt.Printf("strTrimS=%v\n", strTrimS)

	//字串去掉左右指定字符
	strTrim := strings.Trim("a case a", "a")
	fmt.Printf("strTrim=%v\n", strTrim)

	//是否以指定字串開頭
	strHs := strings.HasPrefix("confix", "con")
	fmt.Printf("strHs=%v\n", strHs)

	//是否以指定字串結尾
	strHss := strings.HasSuffix("abc.com", ".com")
	fmt.Printf("strHss=%v\n", strHss)
}

//內建函數*(不用引包)
//package builtin

//1. 取得長度 func len(v Type) int

//2. 處理字串 有中文  []rune(string)

//3. 字串轉整數  func Atoi(s string) (i int, err error)

//4. 整數轉字串  func Itoa(i int) string

//5. 字串轉byte[]    變數=[]byte("hello")

//6. byte[]轉字串    變數=string([]byte(97,98,99))

//7. 10進制轉 2,8,16進制  str=strconv.FormatInt(123,2)

//8. 字串是否包含某字串  func Contains(s, substr string) bool

//9. 計算字串裡有幾個指定的字串 func Count(s, sep string) int

//10. 不區分大小寫的字串比較  func EqualFold(s, t string) bool  區分大小寫 ==

//11. 返回某字串在字串第一次出現的索引值，都沒有返回-1  func Index(s, sep string) int

//12. 返回某字串在字串最後一次出現的索引值，都沒有返回-1  func LastIndex(s, sep string) int

//13. 指定字串取代 func Replace(s, old, new string, n int) string   ,n : 要替換幾個 n<0 則全部替換

//14. 分割字串  func Split(s, sep string) []string

//15. 大小寫 strings.ToLower(str) strings.ToUpper(str)

//16. 字串去左右空格 strings.TrimSpace(str)

//17. 字串去掉左右指定字符 strings.Trim(str)

//18. 字串去掉左 or 右指定字符 strings.TrimLeft(str)

//19. 是否以指定字串開頭  strings.HasPrefix(str)

//20. 是否以指定字串結尾
