package sorts

import (
	"testing"
)

func TestLog2(t *testing.T) {
	v := log2(8)
	if v != 3 {
		t.Errorf("log2(8) want %d, got %d", 3, v)
	}
}

func TestShellSort(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	ShellSort(s)
	ln(s)

	s = []int{2, 1}
	ShellSort(s)
	ln(s)

	s = []int{2}
	ShellSort(s)
	ln(s)

	s = []int{}
	ShellSort(s)
	ln(s)

	s = []int{5, 4, 3, 2, 1}
	ShellSort(s)
	ln(s)

	c := 20
	s = make([]int, c)
	for i, j := c, 0; i > 0; i-- {
		s[j] = i
		j++
	}
	ShellSort(s)
	ln(s)

	for c := 100; c >= 0; c-- {
		s = make([]int, c)
		for i, j := c, 0; i > 0; i-- {
			s[j] = i
			j++
		}
		ShellSort(s)
		ln(s)
	}
}

func BenchmarkShellSort(b *testing.B) {
	c := 100
	for i := 0; i < b.N; i++ {
		s := make([]int, c)
		for i, j := c, 0; i > 0; i-- {
			s[j] = i
			j++
		}

		ShellSort(s)
	}
}
