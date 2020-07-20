package easy167

func twoSum(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1

	for low < high {
		left, right := numbers[low], numbers[high]
		sum := left + right
		if sum > target {
			for low < high && (numbers[high] == right) {
				high--
			}
		} else if sum < target {
			for low < high && (numbers[low] == left) {
				low++
			}
		} else {
			return []int{low + 1, high + 1}
		}
	}

	return []int{-1, -1}
}

func twoSumWithBinary(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		low, high := i+1, len(numbers)-1
		for low <= high {
			mid := low + (high-low)/2
			sum := numbers[i] + numbers[mid]
			if sum > target {
				high = mid - 1
			} else if sum < target {
				low = mid + 1
			} else {
				return []int{i + 1, mid + 1}
			}
		}
	}

	return []int{-1, -1}
}
