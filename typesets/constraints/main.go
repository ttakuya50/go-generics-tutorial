package main

import (
	"constraints"
	"fmt"
)

func main() {
	// intなのでエラーになる
	do1(1)

	// int64なのでエラーにならない
	do1(int64(2))

	// Int64なのでエラーになる
	do1(Int64(3))

	// constraints.Integerを使用しているのでエラーにならない
	do2(Int64(4))
}

type int64OrFloat64 interface {
	int64 | float64
}

func do1[T int64OrFloat64](a T) {
	fmt.Println(a)
}

type Int64 int64

func do2[T constraints.Integer](a T) {
	fmt.Println(a)
}
