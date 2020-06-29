export function lengthOfLongestSubstring(string: string): number {
  let left = 0,
    right = -1,
    answer = 0
  const length = string.length
  const uniqueChars = new Set<string>()

  while (left < length) {
    // 缩小窗口
    if (left !== 0) {
      uniqueChars.delete(string[left - 1])
    }

    // 增大窗口
    while (right + 1 < length && !uniqueChars.has(string[right + 1])) {
      uniqueChars.add(string[right + 1])
      right++
    }

    answer = Math.max(answer, right - left + 1)

    left++
  }

  return answer
}
