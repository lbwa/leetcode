package medium16

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	answer := math.MaxInt32

	for i := 0; i < len(nums); i++ {
		low, high := i+1, len(nums)-1

		for low < high {
			left, right := nums[low], nums[high]
			sum := left + right + nums[i]

			/*
				此处若 sum 为两数之和，那么 sum 需与 target - nums[i] 计算差值，而 answer 应与之前的 nums[i] 计算差值
			*/
			if abs(sum-target) < abs(answer-target) {
				answer = sum
			}

			if sum > target {
				for low < high && (nums[high] == right) {
					high--
				}
			} else if sum < target {
				for low < high && (nums[low] == left) {
					low++
				}
			} else {
				return sum
			}
		}

		for i < len(nums)-1 && (nums[i] == nums[i+1]) {
			i++
		}
	}

	return answer
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
