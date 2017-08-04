//main.go
package main

import (
	"fmt"
	"math/rand"
	// "reflect"
)

type Mahjang struct {
	stack []int
	m_Num int
}

type Player struct {
	cards []int
	tile  [3]uint
	tNum  [3]int
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
	fmt.Println(ply.tile, ply.cards, ply.tNum)
	showa(ply.tile[0])
	showa(ply.tile[1])
	showa(ply.tile[2])
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
	for _, v := range ply.cards {
		t := int(v / 10)
		i := v % 10
		y := ply.tile[t]
		j := y >> (uint(i) * 4) & 15
		k := j<<1 + 1
		ply.tile[t] = y | (k << (uint(i) * 4))
	}
	for i, j := 0, 0; i < 3; i++ {
		y := ply.tile[i]
		for y > 0 {
			if y&1 == 1 {
				j++
			}
			y >>= 1
		}
		ply.tNum[i] = j
		j = 0
	}
}

func (ply *Player) ClearCards() {
	ply.cards = ply.cards[:0]
}

func (ply *Player) grading() int { //level:  AAA/6 ABC/5 JJ/5 AA/4 AB/3 AC/2 J/1 A/0
	dict := make(map[int]int)
	for w := 0; w < 3; w++ {
		x := ply.tile[w]
		x >>= 4
		m := w * 10
		for n := 1; x > 0; n++ {
			i := x & 15
			j := x & 273
			k := m + n
			x >>= 4
			a, b, c := 0, 0, 0
			if i == 1 {
				if j == 273 {
					b = 5
					c = 5
					a = 5
				} else if j == 257 {
					c = 2
					a = 2
				} else if j == 17 {
					b = 3
					a = 3
				} else if j == 1 {
					if n%3 == 2 {
						a = 1
					} else {
						a = 0
					}
				}
			} else if i == 3 {
				if n%3 == 2 {
					a = 5
				} else {
					a = 4
					if j == 273 {
						b = 5
						c = 5
					} else if j == 257 {
						c = 2
					} else if j == 17 {
						b = 3
					}
				}
			} else if i == 7 {
				a = 6
			} else if i == 15 {
				a = 6
				if j == 273 {
					b = 5
					c = 5
				} else if j == 257 {
					c = 2
				} else if j == 17 {
					b = 3
				} else if j == 1 {
					a = 5
				}
			} else {
				continue
			}
			if b > 0 {
				if v, ok := dict[k+1]; ok {
					if b > v {
						dict[k+1] = b
					}
				} else {
					dict[k+1] = b
				}
			}
			if c > 0 {
				if v, ok := dict[k+2]; ok {
					if c > v {
						dict[k+2] = c
					}
				} else {
					dict[k+2] = c
				}
			}
			if v, ok := dict[k]; ok {
				if a > v {
					dict[k] = a
				}
			} else {
				dict[k] = a
			}
		}
	}
	d, e := 0, 0
	for k, v := range dict {
		if d == 0 {
			d, e = k, v
		} else if e > v {
			d, e = k, v
		}
	}
	fmt.Println(dict)
	return d
}

func (ply *Player) ChiPai(v int) {
	t := int(v / 10)
	i := uint(v % 10)
	j := i * 4
	y := ply.tile[t]
	k := ((y >> j & 15) << 1) + 1
	ply.tile[t] = y | k<<j
	ply.tNum[t]++
	ply.cards = append(ply.cards, v)
}

func (ply *Player) DaPai(v int) {
	var b uint = 0
	t := int(v / 10)
	i := uint(v % 10)
	n := i * 4
	y := ply.tile[t]
	if i > 1 {
		b = y & (1<<n - 1)
	}
	j := y >> (n + 1)
	k := j & 7
	y = ((j >> 3 << 4) + k) << n
	ply.tile[t] = y | b
	ply.tNum[t]--
	l, m := len(ply.cards), 0
	for k, p := range ply.cards {
		if m < 1 && v == p {
			m = 1
			continue
		}
		if m > 0 && k > 0 {
			ply.cards[k-1] = p
		}
	}
	ply.cards = ply.cards[:l-1]
}

func (ply *Player) Hand(x uint, n int) bool {
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

func (ply *Player) Hu() bool {
	for i := 0; i < 3; i++ {
		if ply.tNum[i]%3 == 1 {
			// fmt.Println("sole")
			return false
		}

		if !ply.Hand(ply.tile[i], ply.tNum[i]) {
			return false
		}
	}
	return true
}

func (ply *Player) Eat(p int) {
	j := int(p / 10)
	k := p % 10
	x := ply.tile[j]
	x >>= 4
	n := ply.tNum[j]
	fmt.Println(j, k, x, n)
	arr := make([]int, 0, 14)
	jng := make([]int, 0, 7)
	less := make([]int, 0, 7)
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
					less = append(less, b)
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
					less = append(less, b, b)
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
				less = append(less, b, b, b, b)
			}
		} else {
			x >>= 4
			fmt.Println("-")
		}
		b++
	}
	fmt.Println(arr, jng, less)
}

func main() {
	maj := new(Mahjang)
	maj.Init()
	maj.Shuffle()
	maj.Print()
	for i := 0; i < 13; i++ {
		pai := maj.get_pai()
		paylers[0].add_pai(pai)
	}
	paylers[0].FenZu()
	paylers[0].PrintCards()

	// if paylers[0].HuWJang() {
	// 	fmt.Println("Wujiang")
	// }
	// if paylers[0].HuQueSe() {
	// 	fmt.Println("QueSe")
	// }
	// if paylers[0].HuQiDui() {
	// 	fmt.Println("QiDui")
	// }
	// f := paylers[0].HuLiuShunSiXi()
	// fmt.Println(f)
	// h := paylers[0].Hu()
	// if h == true {
	// 	fmt.Println("HU PAI")
	// } else {
	// 	fmt.Println("HU NO")
	// }
	hz := maj.GetNum()
	for i := 1; hz > 1; i++ {
		p := maj.get_pai()
		paylers[0].ChiPai(p)
		paylers[0].Eat(p)
		q := paylers[0].grading()
		paylers[0].DaPai(q)
		hz = maj.GetNum()
		fmt.Printf("Chi %d Da %d ", p, q)
		paylers[0].PrintCards()
	}

	// lst := []int{1, 2, 3, 4, 6, 6}
	// y := Sort(lst)
	// grading2(y)

	// var j, p int
	// for i := 1; hz > 0; i++ {
	// 	p = maj.get_pai()
	// 	hz = maj.GetNum()
	// 	j = i % 4
	// 	q := paylers[j].grading()
	// 	paylers[j].DaPai(q)
	// 	fmt.Printf("hangzhuan %d, payler %d, chipai %d dapai %d  ", hz, j, p, q)
	// 	fmt.Println("...")
	// }
}

func Sort(arr []int) uint {
	var y uint = 0
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

// func showa(y uint64) {
func showa(y uint) {
	for j := 0; y > 0; j++ {
		if y&1 == 1 {
			fmt.Printf("%d", int(j/4))
		}
		y >>= 1
	}
	fmt.Println(" ")
}

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

type MahjangItem struct {
	Tile  int
	Card  int
	Level int
}

func (self MahjangItem) String() string {
	return fmt.Sprintf("<Mahjang(%d, %d)>", self.Tile, self.Card)
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
