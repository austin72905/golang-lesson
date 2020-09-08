package testcase01

import (
	"fmt"
	"testing"
)

//這個框架底層有跑main了

//測試calculate.go裡面的涵式是否正確
func TestAddUpper(t *testing.T) {

	prospectRes := 55
	//調用addUpper
	result := addUpper(10)
	//預計應該出現55 (1+2..+10)
	//如果結果不同代表寫錯了
	if result != prospectRes {
		t.Fatalf("addUpper(10) 執行錯誤，期望值=%v  實際值=%v\n", prospectRes, result)
	}

	//如果正確，輸出日誌
	t.Logf("addUpper(10) 執行正確...")
}

func TestGetSub(t *testing.T) {

	prospectRes := 7
	//調用addUpper
	result := getSub(10, 3)
	//預計應該出現55 (1+2..+10)
	//如果結果不同代表寫錯了
	if result != prospectRes {
		t.Fatalf("getSub(10,3) 執行錯誤，期望值=%v  實際值=%v\n", prospectRes, result)
	}

	//如果正確，輸出日誌
	t.Logf("getSub(10,3) 執行正確...")
}

func TestHello(t *testing.T) {
	fmt.Println("TestHello被調用")
}

//開頭一定要Test
//Test完的函數名要大寫 TestAddUpper

//通过 `go test` 命令，可執行函數  (老師講解是go test -v) 進到testCase01
//func TestXxx(*testing.T)
//其中 Xxx 可以是任何字母数字字符串（但第一个字母不能是 [a-z]）
//将该文件放在与被测试的包相同的包中

//testing框架
//1. 將xxx_test.go 的文件引入
//2. 將Test開頭的函數調用

//細節
//1.一定_test.go結尾

//2. 參數類型是 *testing.T

//3. 測試程序的文件，可以有多個函數

//4. 測試指令
//(1)go test(錯誤才會輸出日誌)
//(2)go test -v (無論錯誤還是正確，都輸出日誌)

//5. 出現錯誤時，可以使用t.Fatalf 來輸出錯誤訊息，並退出程序

//6. t.Logf 可以輸出對應日誌

//7. 測試單一文件 :go test 預測會測試所有_test.go結尾的文件，如果要測試單一個文件，可以指定文件名、要測試函數所在文件
//go test -v calculate_test.go calculate.go

//8. 測試單一函數 :
//go test -v -test.run TestAddUpper (cmder才可用)
