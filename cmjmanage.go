package main

import (
	"fmt"
	"math/rand"
	// "reflect"
)

type stPAI struct {
	m_Type  int //牌类型
	m_Value int //牌字
}

type Card struct {
	Suit   string
	Number int
}

//吃牌顺
type stCHI struct {
	m_Type   int //牌类型
	m_Value1 int //牌字
	m_Value2 int //牌字
	m_Value3 int //牌字
}

//胡牌信息
type stGoodInfo struct {
	m_GoodName  string //胡牌术语
	m_GoodValue int    //胡牌番数
}

type stPAIEx struct {
	m_NewPai stPAI //起的新牌
	m_PaiNum int   //剩余牌数
	m_IsHZ   bool  //是否黄庄
}

type CMJ struct {
	m_MyPAIVec       [6]string  //起的种牌型
	m_ChiPAIVec      [6]string  //吃的种牌型
	m_PengPAIVec     [6]string  //碰的种牌型
	m_GangPAIVec     [6]string  //杠的种牌型
	m_TempChiPAIVec  stCHI      //吃的可选组合
	m_TempPengPAIVec stPAI      //碰的可选组合
	m_TempGangPAIVec stPAI      //杠的可选组合
	m_LastPAI        stPAI      //最后起的牌
	m_GoodInfo       stGoodInfo //胡牌信息

	m_4AK   bool //是否是听四暗刻
	m_9LBD  bool //是否听连宝灯牌型
	m_13Y   bool //是否听十三幺
	m_MKNum int  //明刻数
	m_AKNum int  //暗刻数
}

type MJManage struct {
	m_HZPaiNum int
	// m_MJVec    map
	mMahJ []int
}

type Mahjang struct {
	stack []int
	m_Num int
}

type Player struct {
	cards []int
}

//deck
// var paylers = [4]Player{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
var paylers = [4]Player{}

func (mj *Mahjang) set_Num(num int) {
	mj.m_Num = num
}

func (mj *Mahjang) Init() {
	for i := 1; i < 30; i++ {
		if i%10 != 0 {
			mj.stack = append(mj.stack, i, i, i, i)
		}
	}
	mj.m_Num = len(mj.stack)
}

//shuffle

func (mj *Mahjang) Shuffle() {
	for i := 1; i < len(mj.stack); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			mj.stack[r], mj.stack[i] = mj.stack[i], mj.stack[r]
		}
	}
	mj.m_Num = len(mj.stack)
}

// func (mj *Mahjang) Shuffl2() {
// 	var s []int
// 	r:=rand.Intn(36)

// 	for i := 1; i < len(mj.stack); i++ {
// 		r := rand.Intn(i + 1)
// 		s = append(s, r)
// 	}
// 	fmt.Println(s)
// }

func (mj *Mahjang) Print() {
	fmt.Println(mj.stack)
}

func (mj *Mahjang) get_pai() int {
	mj.m_Num--
	card := mj.stack[mj.m_Num]
	return card
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

func (ply *Player) HuJJang() bool {
	for _, v := range ply.cards {
		if v%10%3 != 2 {
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

func (ply *Player) FenZu() {
	var arr [3][]int
	for _, v := range ply.cards {
		j := int(v / 10)
		k := v % 10
		arr[j] = append(arr[j], k)
	}
	fmt.Println(arr)
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
	maj := new(Mahjang)
	maj.Init()
	maj.Shuffle()
	// maj.Print()

	for i := 0; i < 13; i++ {
		for j := 0; j < 4; j++ {
			pai := maj.get_pai()
			paylers[j].add_pai(pai)
		}
	}

	for j := 0; j < 4; j++ {
		paylers[j].PrintCards()
		if paylers[j].HuWJang() {
			fmt.Println("Wujiang")
		}
		if paylers[j].HuQueSe() {
			fmt.Println("QueSe")
		}
		paylers[j].FenZu()
	}
	// maj.Shuffle()
	maj.Shuffl2()

}

// //deck
// var cards = []Card{{Suit: "Spades", Number: 2}, {Suit: "Spades", Number: 3}, {Suit: "Spades", Number: 4}, {Suit: "Spades", Number: 5}, {Suit: "Spades", Number: 6}, {Suit: "Spades", Number: 7}, {Suit: "Spades", Number: 8}, {Suit: "Spades", Number: 9}, {Suit: "Spades", Number: 10}, {Suit: "Spades", Number: 11}, {Suit: "Spades", Number: 12}, {Suit: "Spades", Number: 13}, {Suit: "Spades", Number: 14}, {Suit: "Hearts", Number: 2}, {Suit: "Hearts", Number: 3}, {Suit: "Hearts", Number: 4}, {Suit: "Hearts", Number: 5}, {Suit: "Hearts", Number: 6}, {Suit: "Hearts", Number: 7}, {Suit: "Hearts", Number: 8}, {Suit: "Hearts", Number: 9}, {Suit: "Hearts", Number: 10}, {Suit: "Hearts", Number: 11}, {Suit: "Hearts", Number: 12}, {Suit: "Hearts", Number: 13}, {Suit: "Hearts", Number: 14}, {Suit: "Diamonds", Number: 2}, {Suit: "Diamonds", Number: 3}, {Suit: "Diamonds", Number: 4}, {Suit: "Diamonds", Number: 5}, {Suit: "Diamonds", Number: 6}, {Suit: "Diamonds", Number: 7}, {Suit: "Diamonds", Number: 8}, {Suit: "Diamonds", Number: 9}, {Suit: "Diamonds", Number: 10}, {Suit: "Diamonds", Number: 11}, {Suit: "Diamonds", Number: 12}, {Suit: "Diamonds", Number: 13}, {Suit: "Diamonds", Number: 14}, {Suit: "Clubs", Number: 2}, {Suit: "Clubs", Number: 3}, {Suit: "Clubs", Number: 4}, {Suit: "Clubs", Number: 5}, {Suit: "Clubs", Number: 6}, {Suit: "Clubs", Number: 7}, {Suit: "Clubs", Number: 8}, {Suit: "Clubs", Number: 9}, {Suit: "Clubs", Number: 10}, {Suit: "Clubs", Number: 11}, {Suit: "Clubs", Number: 12}, {Suit: "Clubs", Number: 13}, {Suit: "Clubs", Number: 14}}

// func Shuffle(slc []Card) {
// 	for i := 1; i < len(slc); i++ {
// 		r := rand.Intn(i + 1)
// 		if i != r {
// 			slc[r], slc[i] = slc[i], slc[r]
// 		}
// 	}
// }

// mutli sorting taken from http://golang.org/pkg/sort/
// multiSorter implements the Sort interface, sorting the cards within.

// type multiSorter struct {
// 	cards []Card
// 	less  []lessFunc
// }

// type lessFunc func(p1, p2 *Card) bool

// func (mj *MJManage) get_pai(slc []Card) Card {
// 	mj.m_HZPaiNum--
// 	card := slc[mj.m_HZPaiNum]
// 	return card
// }
