package easy925

func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	// 在 name 和 typed 上移动两指针
	for j < len(typed) {
		// 正常匹配时，共进
		if i < len(name) && name[i] == typed[j] {
			i++
			j++
			// 否则，存在重复输入时，输入指针前进至非重复字符
		} else if j > 0 && (typed[j] == typed[j-1]) {
			j++
			// 既不匹配，而且不是重复字符时，那么不符合要求
		} else {
			return false
		}
	}
	return i == len(name)
}
