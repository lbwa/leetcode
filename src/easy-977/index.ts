/**
 * 对撞指针，两指针相向而行
 * @param list
 */
export function sortedSquares(list: number[]) {
  const result: number[] = []
  let i = list.length
  let head = 0
  let tail = i - 1
  while (i--) {
    const headItem = list[head] ** 2
    const tailItem = list[tail] ** 2
    if (headItem > tailItem) {
      // 索引是从大到小，故此处赋值 headItem
      result[i] = headItem
      head++
    } else {
      result[i] = tailItem
      tail--
    }
  }
  return result
}
