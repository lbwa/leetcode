package medium18

import "sort"

// FourSum function is used to solve four sum problems.
func FourSum(nums []int, target int) [][]int {
	answer := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		triples := threeSumTarget(nums, i+1, target-nums[i])
		for _, triple := range triples {
			answer = append(answer, append(triple, nums[i]))
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return answer
}

/*
	见 medium15 中对下文方程的解析
*/
func threeSumTarget(nums []int, start int, target int) [][]int {
	answer := [][]int{}

	for i := start; i < len(nums)-1; i++ {
		tuples := twoSumTarget(nums, i+1, target-nums[i])
		for _, tuple := range tuples {
			answer = append(answer, append(tuple, nums[i]))
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return answer
}

/*
	见 medium15 中对下文方程的解析
*/
func twoSumTarget(nums []int, start int, target int) [][]int {
	low, high, answer := start, len(nums)-1, [][]int{}

	for low < high {
		left, right := nums[low], nums[high]
		sum := left + right
		if sum > target {
			for low < high && nums[high] == right {
				high--
			}
		} else if sum < target {
			for low < high && (nums[low] == left) {
				low++
			}
		} else {
			answer = append(answer, []int{left, right})
			for low < high && (nums[high] == right) {
				high--
			}
			for low < high && (nums[low] == left) {
				low++
			}
		}
	}

	return answer
}
