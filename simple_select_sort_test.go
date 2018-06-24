package sorts

import (
	"testing"
)

func TestSimpleSelectionSort(t *testing.T) {
	s := []int{5, 4, 3, 2, 1}
	SimpleSelectionSort(s)
	ln(s)

	s = []int{1, 2, 3, 4, 5}
	SimpleSelectionSort(s)
	ln(s)

	s = []int{4, 5, 3, 2, 1}
	SimpleSelectionSort(s)
	ln(s)

	s = []int{3, 5, 1, 4, 2}
	SimpleSelectionSort(s)
	ln(s)

	s = []int{5, 1} 
	SimpleSelectionSort(s)
	ln(s)

	s = []int{5}
	SimpleSelectionSort(s)
	ln(s)

	s = []int{}
	SimpleSelectionSort(s)
	ln(s)
}

func BenchmarkSimpleSelectionSort(b *testing.B) {
	c := 100
	for i := 0; i < b.N; i++ {
		s := make([]int, c)
		for i, j := c, 0; i > 0; i-- {
			s[j] = i
			j++
		}
		
		SimpleSelectionSort(s)
	}
}
