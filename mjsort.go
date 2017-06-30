//sort.go
package sort

func Sort(arr []uint64) {
	var y uint64 = 0
	for _, i := range arr {
		j := y >> (i * 4) & 15
		k := j<<1 + 1
		y = y | (k << (i * 4))
	}
	for j, k := 0, 0; y > 0; j++ {
		if y&1 == 1 {
			arr[k] = uint64(j / 4)
			k++
		}
		y >>= 1
	}
}
