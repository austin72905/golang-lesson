package model

//Student struct
type Student struct {
	Name  string
	Score float64
}

//如果是小寫
type student2 struct {
	Name  string
	Score float64
}

//如果裡面有個屬性是小寫...
type student3 struct {
	Name  string
	score float64
}

//NewStu2 ...
//用工廠模式(說白了就是用一個函數返回student2 的指針)
func NewStu2(name string, score float64) *student2 {
	//返回一個struct 指針 ，可以用指針變數去接收
	return &student2{
		Name:  name,
		Score: score,
	}
}

//NewStu3 ...
func NewStu3(name string, s float64) *student3 {
	//返回一個struct 指針 ，可以用指針變數去接收
	return &student3{
		Name:  name,
		score: s, //這邊的確可以拿到main呼叫時傳過來的值，但是main包沒辦法把這個屬性拿去用
	}
}

//需要新增一個student3 的方法，回傳score
func (stu student3) GetScore() float64 {
	return stu.score //這樣才拿的到score
}
