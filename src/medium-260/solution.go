package medium260

func singleNumber(nums []int) []int {
	var ans int
	// 结束下文这个迭代后，ans 此时为两个唯一值的异或结果值
	for _, num := range nums {
		ans ^= num
	}
	// 定一个非 0 位作为分组标志
	div := 1
	for (div & ans) == 0 {
		div <<= 1
	}
	// 另外一种找到最低非 0 位的方式是，按位非后与原值进行按位与操作
	// div := ans & (-ans)
	var a, b int
	// 分组异或，分别在两个部分异或，进而得到在子数组中的唯一值
	for _, num := range nums {
		if (div & num) == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	// 综合来看，此法通过分治的思维将问题降维至两个子数组中的唯一值问题

	return []int{a, b}
}
