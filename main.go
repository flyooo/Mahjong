//main.go
package main

import (
	"fmt"
	"math/rand"
	// "sort"
	// "reflect"
)

type Mahjang struct {
	stack []int
	m_Num int
}

type Player struct {
	cards []int
	tile  [3][]int
}

var paylers = [4]Player{}

func (mj *Mahjang) SetNum(num int) {
	mj.m_Num = num
}

func (mj *Mahjang) GetNum() int {
	return mj.m_Num
}

func (mj *Mahjang) Init() {
	for i := 1; i < 30; i++ {
		if i%10 != 0 {
			mj.stack = append(mj.stack, i, i, i, i)
		}
	}
}

//shuffle
func (mj *Mahjang) Shuffle() {
	for i := 1; i < len(mj.stack); i++ {
		r := rand.Intn(i + 1)
		mj.stack[r], mj.stack[i] = mj.stack[i], mj.stack[r]
	}
	mj.m_Num = len(mj.stack)
}

func (mj *Mahjang) Shuffl2() {
	n := len(mj.stack)
	for i := 1; i < n; i++ {
		r := rand.Intn(n)
		if i != r {
			mj.stack[r], mj.stack[i] = mj.stack[i], mj.stack[r]
		}
	}
	mj.m_Num = len(mj.stack)
}

func (mj *Mahjang) Print() {
	fmt.Println(mj.stack)
}

func (mj *Mahjang) get_pai() int {
	mj.m_Num--
	card := mj.stack[mj.m_Num]
	return card
}

func (mj *Mahjang) deal(ply []Player) {
	for i := 0; i < 13; i++ {
		for j := 0; j < 4; j++ {
			pai := mj.get_pai()
			fmt.Println(pai)
			ply[i].add_pai(pai)
		}
	}
}

func (ply *Player) add_pai(pai int) {
	ply.cards = append(ply.cards, pai)
}

func (ply *Player) PrintCards() {
	fmt.Println(ply.cards)
}

func (ply *Player) HuWJang() bool {
	for _, v := range ply.cards {
		if v%10%3 == 2 {
			return false
		}
	}
	return true
}

func (ply *Player) HuQueSe() bool {
	var i = [3]int{}
	for _, v := range ply.cards {
		j := int(v / 10)
		if j == 0 {
			i[0]++
		} else if j == 1 {
			i[1]++
		} else if j == 2 {
			i[2]++
		}
	}
	for _, vv := range i {
		if vv == 0 {
			return true
		}
	}
	return false
}

func (ply *Player) HuLiuShunSiXi() int {
	fan, liu, sxi := 0, 0, 0
	dict := make(map[int]int)
	for _, c := range ply.cards {
		if i, ok := dict[c]; ok {
			dict[c] = i + 1
		} else {
			dict[c] = 1
		}
	}
	for _, v := range dict {
		if v == 4 {
			fan++
			sxi++
		} else if v == 3 {
			liu++
		}
	}
	if liu == 4 {
		fan += 2
	} else if liu >= 2 {
		fan++
	}
	fmt.Println(ply.cards)
	fmt.Println(dict, liu, sxi)
	return fan
}

func (ply *Player) HuJJang() bool {
	for _, v := range ply.cards {
		if v%10%3 != 2 {
			return false
		}
	}
	return true
}

func (ply *Player) HuQiDui() bool {
	dict := make(map[int]int)
	for _, c := range ply.cards {
		if i, ok := dict[c]; ok {
			dict[c] = i + 1
		} else {
			dict[c] = 1
		}
	}

	if len(dict) == 7 {
		return true
	} else if len(dict) > 7 {
		return false
	} else {
		for _, v := range dict {
			if v == 1 || v == 3 {
				return false
			}
		}
		return true
	}
}

func (ply *Player) FenZu() {
	ply.tile[0] = ply.tile[0][:0]
	ply.tile[1] = ply.tile[1][:0]
	ply.tile[2] = ply.tile[2][:0]
	for _, v := range ply.cards {
		j := int(v / 10)
		k := v % 10
		ply.tile[j] = append(ply.tile[j], k)
	}
	fmt.Println(ply.tile)
}

func (ply *Player) ClearCards() {
	ply.cards = ply.cards[:0]
}

func (ply *Player) grading() { //level: AAAA/10 AAA/5  AA/3 AB/2 AC/1 A/0
	l, n := 0, 0
	for o := 0; o < 3; o++ {
		l += len(ply.tile[o])
	}
	a := make(MahjangSlice, l)
	for w := 0; w < 3; w++ {
		lst := ply.tile[w]
		x := Sort(lst)
		x >>= 4
		dict := make(map[int]int)
		for c := 1; x > 0; c++ {
			i := x & 15
			j := x & 273
			x >>= 4
			if i > 0 {
				k := 0
				if j == 273 {
					k = 3
					if v, ok := dict[c+2]; ok {
						dict[c+2] = v + 1
					} else {
						dict[c+2] = 1
					}
					if v, ok := dict[c+1]; ok {
						dict[c+1] = v + 2
					} else {
						dict[c+1] = 2
					}
				} else if j == 257 {
					k = 1
					if v, ok := dict[c+2]; ok {
						dict[c+2] = v + 1
					} else {
						dict[c+2] = 1
					}
				} else if j == 17 {
					k = 2
					if v, ok := dict[c+1]; ok {
						dict[c+1] = v + 2
					} else {
						dict[c+1] = 2
					}
				} else if j == 1 {
					k = 0
				}

				// i == 1
				if i == 3 {
					k += 3
				} else if i == 7 {
					k += 5
				} else if i == 15 {
					k += 10
				}

				if v, ok := dict[c]; ok {
					dict[c] = v + k
				} else {
					dict[c] = k
				}
			}
		}
		for k, v := range dict {
			a[n] = &MahjangItem{w, k, v}
			n++
		}
	}
	//sort.Sort(a)
	fmt.Println(a)
}

func (ply *Player) Eat(p int) {
	j := int(p / 10)
	k := p % 10
	lst := ply.tile[j]
	lst = append(lst, k)
	x := Sort(lst)
	n := len(lst)
	fmt.Println(j, k, lst)
	arr := make([]int, 0, 14)
	jng := make([]int, 0, 7)
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
					// return false
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
					// return false
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
				// return false
			}
		} else {
			x >>= 4
			fmt.Println("-")
		}
		b++
	}
	fmt.Println(arr)
	fmt.Println(jng)

}

func main() {
	maj := new(Mahjang)
	maj.Init()
	maj.Shuffle()
	for n := 0; n < 1; n++ {
		maj.Shuffl2()
		for j := 0; j < 4; j++ {
			paylers[j].ClearCards()
		}
		for i := 0; i < 13; i++ {
			for j := 0; j < 4; j++ {
				pai := maj.get_pai()
				paylers[j].add_pai(pai)
			}
		}

		for j := 0; j < 4; j++ {
			paylers[j].FenZu()
			// // paylers[j].PrintCards()
			// if paylers[j].HuWJang() {
			// 	fmt.Println("Wujiang")
			// }
			// if paylers[j].HuQueSe() {
			// 	fmt.Println("QueSe")
			// }
			// if paylers[j].HuQiDui() {
			// 	fmt.Println("QiDui")
			// }
			// f := paylers[j].HuLiuShunSiXi()
			// fmt.Println(f)
		}
	}
	hz := maj.GetNum()
	var j, p int
	for i := 1; hz > 0; i++ {
		p = maj.get_pai()
		hz = maj.GetNum()
		j = i % 4
		fmt.Println(p)
		// paylers[j].Eat(p)
		// fmt.Println(hz, p, j)
		paylers[j].grading()
		// if paylers[j].ting {
		// 	paylers[j].checkhand()
		// } else {
		// 	paylers[j].checkting()
		// }
	}

	// a := make(GameDownloadSlice, 7)
	// for i := 0; i < len(a); i++ {
	// 	a[i] = &GameDownloadItem{i + 1, rand.Intn(1000)}
	// }

	// fmt.Println(a)
	// sort.Sort(a)
}

func Sort(arr []int) uint64 {
	var y uint64 = 0
	for _, i := range arr {
		j := y >> (uint(i) * 4) & 15
		k := j<<1 + 1
		y = y | (k << (uint(i) * 4))
	}
	for x, j, k := y, 0, 0; x > 0; j++ {
		if x&1 == 1 {
			arr[k] = j / 4
			k++
		}
		x >>= 1
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

func Hand(x uint64, n int) bool {
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

type MahjangItem struct {
	Tile  int
	Card  int
	Level int
}

func (self MahjangItem) String() string {
	return fmt.Sprintf("<Mahjang(%d-%d,%d)>", self.Tile, self.Card, self.Level)
}

type MahjangSlice []*MahjangItem

func (p MahjangSlice) Len() int {
	return len(p)
}

func (p MahjangSlice) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

// 根据游戏下载量 降序 排列
func (p MahjangSlice) Less(i int, j int) bool {
	return p[i].Level > p[j].Level
}
