package sorts

import (
	"github.com/ZhangGuangxu/search"
)

type QuickSort struct {
	s []int
	tLen int
	t []int
	bst *search.BinarySearchTree
}

func NewQuickSort(s []int) *QuickSort {
	return &QuickSort{
		s: s,
		t: make([]int, len(s)),
		bst: search.NewBinarySearchTree(),
	}
}

func (q *QuickSort) Sort() {
	for a := range q.s {
		q.insert(a)
	}
}

func (q *QuickSort) insert(a int) {
	q.bst.Insert(a)
}

func (q *QuickSort) Result() ([]int, error) {
	return q.bst.InorderTraverse()
}
