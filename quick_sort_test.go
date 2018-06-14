package sorts

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	s := make([]int, 0, 100)
	for i, j := 49, 50; i >= 0 && j < 100; {
		s = append(s, i)
		s = append(s, j)
		i--
		j++
	}
	q := NewQuickSort(s)
	q.Sort()
	r, err := q.Result()
	ln(r, err)
	for i := 0; i < len(r)-1; i++ {
		if r[i] >= r[i+1] {
			t.Error("sort result is wrong")
		}
	}
}
