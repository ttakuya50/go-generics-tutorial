package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func main() {
	// ソートする
	a := []int64{1, 4, 2, 5, 3}
	slices.Sort(a)
	fmt.Printf("a=%v\n", a)

	// ソート済みか判定
	b := []int64{1, 4, 2, 5, 3}
	ok1 := slices.IsSorted(b)
	fmt.Printf("ok1=%v\n", ok1)
	ok2 := slices.IsSorted(a)
	fmt.Printf("ok2=%v\n", ok2)

	// 含まれているか判定
	c := []int{1, 2, 3, 4, 5}
	ok3 := slices.Contains(c, 5)
	ok4 := slices.Contains(c, 10)
	fmt.Printf("ook3=%v\n", ok3)
	fmt.Printf("ok4=%v\n", ok4)

	d := []string{"apple", "banana", "grape", "orange"}
	ok5 := slices.Contains(d, "grape")
	fmt.Printf("ok5=%v\n", ok5)
}
