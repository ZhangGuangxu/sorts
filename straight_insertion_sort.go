package sorts

// StraightInsertionSort do straight insertion sort
func StraightInsertionSort(s []int) {
	sLen := len(s)
	for i := 1; i < sLen; i++ {
		if s[i] < s[i-1] {
			tmp := s[i]
			var j int
			for j = i-1; j >= 0 && s[j] > tmp; j-- {
				s[j+1] = s[j]
			}
			s[j+1] = tmp
		}
	}
}
