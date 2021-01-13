package medium189

func rotate(nums []int, k int) {
	extra := make([]int, len(nums))
	for i, v := range nums {
		extra[(i+k)%len(nums)] = v
	}
	copy(nums, extra)
}

/*
	nums = "----->-->"; k =3
	result = "-->----->";

	reverse "----->-->" we can get "<--<-----"
	reverse "<--" we can get "--><-----"
	reverse "<-----" we can get "-->----->"
*/
func rotate0(nums []int, k int) {
	var reverse func([]int)
	reverse = func(nums []int) {
		n := len(nums)
		for i := 0; i < n/2; i++ {
			nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
		}
	}

	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
