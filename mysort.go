package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type GameDownloadItem struct {
	GameID        int // 游戏ID
	DownloadTimes int // 下载次数
}

func (self GameDownloadItem) String() string {
	return fmt.Sprintf("<Item(%d, %d)>", self.GameID, self.DownloadTimes)
}

type GameDownloadSlice []*GameDownloadItem

func (p GameDownloadSlice) Len() int {
	return len(p)
}

func (p GameDownloadSlice) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

// 根据游戏下载量 降序 排列
func (p GameDownloadSlice) Less(i int, j int) bool {
	return p[i].DownloadTimes > p[j].DownloadTimes
}

func main() {
	a := make(GameDownloadSlice, 7)
	for i := 0; i < len(a); i++ {
		a[i] = &GameDownloadItem{i + 1, rand.Intn(1000)}
	}

	fmt.Println(a)
	sort.Sort(a)
	fmt.Println(a)
}
