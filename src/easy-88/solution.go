package easy88

func merge(nums1 []int, m int, nums2 []int, n int) {
	p, q, tail := m-1, n-1, m+n-1
	for ; p >= 0 || q >= 0; tail-- {
		var current int
		if p == -1 {
			current = nums2[q]
			q--
		} else if q == -1 {
			current = nums1[p]
			p--
		} else if nums1[p] > nums2[q] {
			// 因为是由后向前插入，故此处选择较大值
			current = nums1[p]
			p--
		} else {
			current = nums2[q]
			q--
		}

		nums1[tail] = current
	}
}
