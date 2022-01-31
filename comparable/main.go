package main

import "log"

func main() {
	// 数値の場合
	a := []int64{1, 2, 3, 4, 5}
	res1 := find(a, int64(3))
	log.Printf("res1=%v\n", res1)

	// 文字列の場合
	b := []string{"apple", "banana", "grape", "orange"}
	res2 := find(b, "orange")
	log.Printf("res2=%v\n", res2)
}

// find 引数の比較をして同じ値があればindexを返却
func find[T comparable](a []T, b T) int {
	for i, v := range a {
		if v == b {
			return i
		}
	}
	return -1
}
