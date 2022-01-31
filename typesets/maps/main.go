package main

import (
	"fmt"
	"golang.org/x/exp/maps"
)

func main() {
	// mapのvalue部分だけ抽出
	m := map[string]int{"udon": 500, "ra-men": 800, "soba": 700}
	mv := maps.Values(m)
	for _, v := range mv {
		fmt.Printf("v=%v\n", v)
	}

	// mapのkey部分だけ抽出
	mk := maps.Keys(m)
	for _, k := range mk {
		fmt.Printf("k=%v\n", k)
	}

	a := map[string]int{"banana": 500, "apple": 300}
	b := map[string]int{"banana": 500, "apple": 300}
	ok1 := maps.Equal(a, b)
	fmt.Printf("ok1=%v\n", ok1)

	c := map[string]int{"banana": 500, "apple": 400}
	ok2 := maps.Equal(b, c)
	fmt.Printf("ok2=%v\n", ok2)
}
