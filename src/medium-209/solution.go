package medium209

import "math"

func minSubArrayLen(s int, nums []int) int {
	slow, fast, answer, sum := 0, 0, math.MaxInt32, 0

	for fast < len(nums) {
		sum += nums[fast]
		fast++

		for sum >= s {
			answer = min(fast-slow, answer)
			deleted := nums[slow]
			slow++
			sum -= deleted
		}
	}

	if answer == math.MaxInt32 {
		return 0
	}

	return answer
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
