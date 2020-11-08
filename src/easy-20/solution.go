package easy20

func isValid(s string) bool {
	charMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack []rune
	for _, r := range s {
		// 如果是开始的括号时，加入到栈中，待匹配
		if _, ok := charMap[r]; ok {
			stack = append(stack, r)
			continue
		}

		// 没有待匹配括号时，但需要匹配时，则不满足题目要求，退出执行
		if len(stack) < 1 {
			return false
		}

		// 否则取栈顶待匹配括号来匹配当前括号
		topIndex := len(stack) - 1
		top := stack[topIndex]
		if charMap[top] == r {
			stack = stack[:topIndex]
			continue
		}

		// 当栈顶括号无法匹配时，即不满足题目要求
		return false
	}
	return len(stack) == 0
}
