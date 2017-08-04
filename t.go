package main

import (
	"fmt"
)

func deletei(i int, arr []int) []int {
	t, n := 0, len(arr)
	l := make([]int, 0, n)
	for _, v := range arr {
		if t < 1 && v == i {
			t = 1
			continue
		}
		l = append(l, v)
	}
	return l
}

func main() {
	// s := []int{11, 22, 33, 44, 55, 33, 66}
	// l := deletei(33, s)
	// fmt.Println(l)

	// j := 1 << uint(2*4)
	// fmt.Println(j)
	i := 26
	j := int(i / 10)
	fmt.Println(j)
}
