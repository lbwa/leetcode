package easy344

func reverseString(s []byte) {
	low, high := 0, len(s)-1
	for low < high {
		s[low], s[high] = s[high], s[low]
		low++
		high--
	}
}
