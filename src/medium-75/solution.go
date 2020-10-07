package medium75

import "sort"

const (
	red = iota
	white
	blue
)

func sortColors0(nums []int) {
	sort.Ints(nums)
}

// Option 1, singly pointer
func sortColors1(nums []int) {
	swapColors := func(colors []int, target int) (countTarget int) {
		for i, color := range colors {
			if color == target {
				colors[i], colors[countTarget] = colors[countTarget], colors[i]
				countTarget++
			}
		}
		return
	}
	// 一次遍历，将所有 0 前置
	count0 := swapColors(nums, red)
	// 一次遍历，将所有 1 前置，但排在所有 0 之后
	swapColors(nums[count0:], white)
}

// Option ,2, two pointers
func sortColors2(nums []int) {
	p0, p1 := 0, 0
	for i, color := range nums {
		if color == red {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if color == white {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}

func sortColors3(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for ; i <= p2 && nums[i] == blue; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		if nums[i] == red {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
