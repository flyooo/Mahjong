package main

import (
	"fmt"
	// "strconv"
)

func bin(v uint64) {
	var s [14]uint64
	for i, j := v, 13; i > 0; j-- {
		s[j] = i & 1
		i = i >> 1
	}
	fmt.Println(s)

}

func sorta(arr []uint64) uint64 {
	var y uint64 = 0
	for _, i := range arr {
		j := y >> (i * 4) & 15
		k := j<<1 + 1
		y = y | (k << (i * 4))
	}
	return y
}

func sortb(v uint64) uint64 {
	var x uint64
	for x = 0; v > 0; {
		x = ((x << 4) + (v & 15))
		v >>= 4
	}
	return x
}

func showa(y uint64) {
	for j := 0; y > 0; j++ {
		if y&1 == 1 {
			fmt.Printf("%d", int(j/4))
		}
		y >>= 1
	}
	fmt.Println(" ")
}

func hand(x uint64, n int) bool {
	arr := make([]int, 0, 14)
	jng := make([]int, 0, 7)
	if n%3 == 1 {
		fmt.Println("sole")
		return false
	}
	for a, b := 0, 1; x > 0; {
		i := x & 15
		if i == 1 {
			if (x & 273) == 273 { // 'abc'
				x >>= 4
				j := (x & 15) >> 1
				k := (x >> 5) & 7
				x = (((x >> 8 << 4) + k) << 4) + j
				arr = append(arr, b, b+1, b+2)
				fmt.Println(b, b+1, b+2)
			} else {
				if n%3 == 2 && a != 0 { //rollback
					r := arr[len(arr)-1]
					if r == b {
						fmt.Println(r, b, x)
						x = x | 3
					} else if r > b {
						j := ((x >> 3) & 15) + 1
						x = (((x >> 8 << 4) + j) << 4) + 1
					} else if r < b {
						x = (x << 4) + 1
						b--
					}
					a = 0
					arr = arr[:len(arr)-1]
					fmt.Println("rollback ", b, x)
					continue //jump b++
				} else {
					x >>= 4
					fmt.Println("ABC false")
					return false
				}
			}

		} else if i == 3 { //aabbccdd, aabbcc,aa
			if n%3 == 0 {
				if (x & 819) == 819 { //'aabbcc'
					x >>= 4
					j := (x & 15) >> 2
					k := (x >> 6) & 3
					x = (((x >> 8 << 4) + k) << 4) + j
					arr = append(arr, b, b+1, b+2, b, b+1, b+2)
					fmt.Println(b, b+1, b+2)
				} else {
					x >>= 4
					fmt.Println("aa false")
					return false
				}
			} else {
				if (x & 819) == 819 { //'aabbcc' of 'aabbccdd'
					x >>= 4
					j := (x & 15) >> 2
					k := (x >> 6) & 3
					x = (((x >> 8 << 4) + k) << 4) + j
					arr = append(arr, b, b+1, b+2, b, b+1, b+2)
					fmt.Println(b, b+1, b+2)
				} else {
					x >>= 4
					jng = append(jng, b)
					continue
				}
			}
		} else if i == 7 { //aaabcd, aaabc, aaa
			if (x & 4375) == 279 { //'aaabc_'
				x >>= 4
				j := (x & 15) >> 1
				k := (x >> 5) & 7
				x = (((x >> 8 << 4) + k) << 4) + j
				arr = append(arr, b, b+1, b+2)
				fmt.Println(b, b+1, b+2)
				jng = append(jng, b)
			} else if (x & 4375) == 4375 {
				x >>= 4
				a = b
				arr = append(arr, b, b, b)
				fmt.Println(b, b, b)
			} else {
				x >>= 4
				arr = append(arr, b, b, b)
				fmt.Println(b, b, b)
			}
		} else if i == 15 { //aaaabbbbcccc, aaaabbcc, aaaabc

			if (x & 4095) == 4095 { //'aaaabbbbcccc' ''
				x >>= 12
				arr = append(arr, b, b, b, b+1, b+1, b+1, b+2, b+2, b+2, b, b+1, b+2)
				fmt.Println(b, b+1, b+2)
				b = b + 2
			} else if (x & 831) == 831 { //'aaaabbcc'
				x >>= 4
				j := (x & 15) >> 2
				k := (x >> 6) & 3
				x = (((x >> 8 << 4) + k) << 4) + j
				jng = append(jng, b)
				arr = append(arr, b, b+1, b+2, b, b+1, b+2)
				fmt.Println(b, b+1, b+2)
			} else if (x & 287) == 287 { //'aaaabc  100010001111'
				x >>= 4
				j := (x & 15) >> 1
				k := (x >> 5) & 7
				x = (((x >> 8 << 4) + k) << 4) + j
				arr = append(arr, b, b, b, b, b+1, b+2)
				fmt.Println(b, b, b, b+1, b+2)
			} else {
				x >>= 4
				fmt.Println("aaaa false")
				return false
			}
		} else {
			x >>= 4
			fmt.Println("-")
		}
		b++
	}
	fmt.Println(arr)
	fmt.Println(jng)
	return true
}

func main() {
	// var arr = []uint64{6, 5, 5, 6, 7, 7, 5, 5}
	// x := sorta(arr)
	// fmt.Println(strconv.FormatUint(x, 2))
	// showa(x)
	// x >>= 4
	// handu(x, len(arr))

	// fmt.Println("---------------------------------")

	// var lst = []uint64{3, 5, 6, 6, 5, 9, 7, 7, 4, 8, 4, 2, 2, 2}
	// y := sorta(lst)
	// showa(y)
	// z := sortb(y)

	// y >>= 4
	// fmt.Println(strconv.FormatUint(y, 2))
	// fmt.Println(strconv.FormatUint(z, 2))
	// // handu(y, len(lst))
	// hu(y, len(lst))
	// hu(z, len(lst))

	// var lst = []uint64{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 5, 5}
	// var lst = []uint64{2, 2, 2, 2, 3, 3, 4, 4, 5, 6, 7}
	// var lst = []uint64{2, 2, 2, 2, 3, 4, 5, 6, 7}
	// var lst = []uint64{2, 2, 2, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 9}
	var lst = []uint64{2, 2, 2, 3, 4, 5, 6, 7, 7, 8, 9}
	// var lst = []uint64{2, 2, 2, 3, 4, 7, 8, 9}
	// var lst = []uint64{2, 2, 2, 3, 4, 5, 8, 6} //assert 'abc false'
	// var lst = []uint64{1, 2, 2, 2, 3, 4, 5, 6}
	// var lst = []uint64{2, 2, 3, 3, 4, 4, 5, 5}

	y := sorta(lst)
	showa(y)
	// z := sortb(y)
	y >>= 4
	logic := hand(y, len(lst))
	fmt.Println(logic)
}
