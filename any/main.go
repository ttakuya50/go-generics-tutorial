package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	a := []int64{1, 2, 3, 4, 5}
	res1 := mapFunc(a, func(t int64) string {
		return "<" + fmt.Sprint(t) + ">"
	})
	log.Printf("res1=%v\n", res1)

	b := []int{1, 2, 3, 4, 5}
	res2 := mapFunc(b, func(t int) string {
		return "<" + fmt.Sprint(t) + ">"
	})
	log.Printf("res2=%v\n", res2)

	c := []string{"1", "2", "3", "4", "5"}
	res3 := mapFunc(c, func(t string) int {
		ato, _ := strconv.Atoi(t)
		return ato
	})
	log.Printf("res3=%v\n", res3)
}

func mapFunc[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a), cap(a))
	for i, v := range a {
		n[i] = f(v)
	}
	return n
}
