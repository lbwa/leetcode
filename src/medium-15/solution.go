package medium15

import "sort"

// ThreeSum is used to
func ThreeSum(nums []int) [][]int {
	answer := [][]int{}

	// 将 slice 排序
	sort.Ints(nums)

	// 将三数之和问题转换为固定一项的两数之和问题
	for i := 0; i < len(nums); i++ {
		tuples := twoSumTarget(nums, i+1, 0-nums[i])
		for _, tuple := range tuples {
			answer = append(answer, append(tuple, nums[i]))
		}
		// 通过以下迭代来防止三项中的首项出现重复
		for i < len(nums)-1 && (nums[i] == nums[i+1]) {
			i++
		}
	}

	return answer
}

func twoSumTarget(nums []int, start int, target int) [][]int {
	low, high, answer := start, len(nums)-1, [][]int{}

	for low < high {
		left, right := nums[low], nums[high]
		sum := left + right
		if sum > target {
			for low < high && (nums[high] == right) {
				high--
			}
		} else if sum < target {
			for low < high && (nums[low] == left) {
				low++
			}
		} else {
			answer = append(answer, []int{left, right})
			// 通过迭代跳过重复项
			for low < high && (nums[low] == left) {
				low++
			}
			for low < high && (nums[high] == right) {
				high--
			}
		}
	}

	return answer
}
