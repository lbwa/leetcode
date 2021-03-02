package medium338

func countBits(num int) []int {
	var onesCount func(int) int
	onesCount = func(x int) (ones int) {
		for x > 0 {
			x &= (x - 1) // 将二进制最右(低位)的 1 变为 0
			ones++
		}
		return ones
	}

	bits := make([]int, num+1)
	for i := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}
