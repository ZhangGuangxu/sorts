package sorts

import (
	// "math"
	"fmt"
)

func log2(x int) int {
	if x < 1 {
		panic(fmt.Errorf("log2(x), x=%v", x))
	}

	var r int
	for {
		x >>= 1
		if x < 1 {
			break
		}
		r++
	}
	return r
}

// ShellSort do shell sort
func ShellSort(s []int) {
	sLen := len(s)
	if sLen < 2 {
		return
	}

	//t := int(math.Floor(math.Log2(float64(sLen+1))))
	t := log2(sLen)
	interval := sLen
	for i := 0; i < t; i++ {
		//interval = int(math.Exp2(float64(t-i+1)) - 1)
		interval = 1 << uint(t-i) - 1
		for j := interval; j < sLen; j++ {
			if s[j] < s[j-interval] {
				tmp := s[j]
				var k int
				for k = j-interval; k >= 0 && tmp < s[k]; k -= interval {
					s[k+interval] = s[k]
				}
				s[k+interval] = tmp
			}
		}
	}
}
