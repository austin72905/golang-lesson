package testcase02

import "testing"

func TestStore(t *testing.T) {
	//new 一個 monster
	mons := Monster{
		Name:  "zaza",
		Age:   10,
		Skill: "射精",
	}
	//預期結果
	prospectRes := true
	//實際結果
	result := mons.Store() //底層幫我們優化過了 這樣寫跟&mons一樣，只要看method是不是定義指針
	if !result {
		t.Fatalf("mons.Store() 錯誤 ， 期望值 =%v  實際值 =%v", prospectRes, result)
	}
	t.Logf("mons.Store() 執行成功")

}

func TestRestore(t *testing.T) {

	mons := Monster{}

	//預期結果
	prospectRes := true
	//實際結果
	result := mons.Retore()
	if !result {
		t.Fatalf("mons.Restore() 錯誤 ， 期望值 =%v  實際值 =%v", prospectRes, result)
	}

	//判斷反序列化的屬性值是否正確
	if mons.Name != "zaza" {
		t.Fatalf("mons.Restore() 錯誤 ， 期望值 =%v  實際值 =%v", "zaza", result)
	}

	t.Logf("mons.Restore() 執行成功")
}

//寫兩個函數
//1. 序列化struct 並將json字串寫入文件
//2. 反序列化json字串並取得struct屬性值
//3. 測試
