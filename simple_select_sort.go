package sorts

// SimpleSelectionSort do simple select sort
func SimpleSelectionSort(s []int) {
	for i := len(s); i > 1; i-- {
		maxIdx := 0
		for j := 1; j < i; j++ {
			if s[j] > s[maxIdx] {
				maxIdx = j
			}
		}
		if maxIdx != i-1 {
			s[maxIdx], s[i-1] = s[i-1], s[maxIdx]
		}
	}
}
