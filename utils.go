package sorts

import (
	"log"
)

var ln = log.Println

// a quick check
func isOrdered(s []int) bool {
	iEnd := len(s) - 1
	for i := 0; i < iEnd; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}
