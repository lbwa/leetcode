package easy191

func hammingWeight(num uint32) int {
	var ans int
	for num > 0 {
		num &= (num - 1)
		ans++
	}
	return ans
}
