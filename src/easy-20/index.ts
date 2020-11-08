const charMap: Record<string, string> = {
  '(': ')',
  '{': '}',
  '[': ']'
}

/**
 * 思路：
 * 1. 以栈结构存储匹配缓存
 * 2. 当栈顶匹配到收括号时，栈顶退出栈
 * 3. 当迭代所有字符后，若 stack 栈为空，那么表示所有括号得到匹配
 * @param string 待测字符串
 */
export function isValid(string: string) {
  if (string.length < 1) return true

  const stack = []

  for (const char of string) {
    // 为收括号时
    if (charMap[char]) {
      stack.push(char)
      continue
    }
    // 取栈顶括号的收括号匹配当前 char
    if (charMap[stack[stack.length - 1]] === char) {
      stack.pop()
      continue
    }
    // 与栈顶的括号不匹配，跳出判断，返回 false
    return false
  }

  return stack.length === 0
}
