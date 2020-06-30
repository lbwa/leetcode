export function minSubArrayLen(target: number, numbers: number[]) {
  const length = numbers.length

  if (length < 1) return 0

  let left = 0,
    right = 0,
    sum = 0,
    answer = Infinity
  // 增大窗口
  while (right < length) {
    sum += numbers[right]

    // 缩小窗口
    while (sum >= target) {
      // 仅当左右边界都改变时，才会更新 answer 变量
      answer = Math.min(answer, right - left + 1)
      sum -= numbers[left]
      left++
    }

    right++
  }

  return answer === Infinity ? 0 : answer
}
