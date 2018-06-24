package sorts

import (
	"testing"
)

func TestStraightInsertionSort(t *testing.T) {
	s := []int{5, 4, 3, 2, 1} 
	StraightInsertionSort(s)
	ln(s)

	s = []int{1, 2, 3, 4, 5} 
	StraightInsertionSort(s)
	ln(s)

	s = []int{3, 4, 2, 1, 5} 
	StraightInsertionSort(s)
	ln(s)

	s = []int{5, 1, 2, 4, 3} 
	StraightInsertionSort(s)
	ln(s)

	s = []int{5, 1} 
	StraightInsertionSort(s)
	ln(s)

	s = []int{5}
	StraightInsertionSort(s)
	ln(s)

	s = []int{}
	StraightInsertionSort(s)
	ln(s)
}

func BenchmarkStraightInsertionSort(b *testing.B) {
	c := 100
	for i := 0; i < b.N; i++ {
		s := make([]int, c)
		for i, j := c, 0; i > 0; i-- {
			s[j] = i
			j++
		}

		StraightInsertionSort(s)
	}
}
