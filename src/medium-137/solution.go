package medium137

func singleNumber(nums []int) int {
	seenOnce, seenTwice := 0, 0

	for _, num := range nums {
		// golang 中使用前缀 ^ 表示按位非，而不是前缀 ~
		seenOnce = ^seenTwice & (seenOnce ^ num)
		seenTwice = ^seenOnce & (seenTwice ^ num)
	}

	return seenOnce
}
