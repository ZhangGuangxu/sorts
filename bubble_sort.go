package sorts

// type Sortable interface {
// 	func Len() int
// 	func Less(i, j int) bool
// 	func Swap(i, j int)
// }

// BubbleSort do bubble sort
func BubbleSort(s []int) []int {
	for i := len(s)-1; i > 0; i-- {
		swap := false
		for j := 0; j < i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	return s
}
