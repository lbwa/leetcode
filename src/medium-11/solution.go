package medium11

func maxArea(height []int) int {
	left, right, answer := 0, len(height)-1, 0
	for right > left {
		area := (right - left) * min(height[left], height[right])
		answer = max(area, answer)
		if height[left] >= height[right] {
			right--
		} else {
			left++
		}
	}
	return answer
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
