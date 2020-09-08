package testcase01

func addUpper(n int) int {
	res := 0
	//這邊故意寫錯
	for i := 0; i <= n-1; i++ {
		res += i
	}
	return res
}

func getSub(n1 int, n2 int) int {
	return n1 - n2
}
