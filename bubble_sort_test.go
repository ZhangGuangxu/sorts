package sorts

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	s := []int{10, 11, 12, 13, 9, 8, 2, 1, 5, 21, 19, 18, 17, 16, 15}
	ln(BubbleSort(s)) // [1 2 5 8 9 10 11 12 13 15 16 17 18 19 21]
}

func BenchmarkBubbleSort(b *testing.B) {
	c := 100
	for i := 0; i < b.N; i++ {
		s := make([]int, c)
		for i, j := c, 0; i > 0; i-- {
			s[j] = i
			j++
		}

		BubbleSort(s)
	}
}
