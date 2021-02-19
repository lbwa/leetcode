/**
 * @ref https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Expressions_and_Operators#bitwise_operators
 */
export function singleNumber(nums: number[]): number {
  let seenOnce = 0,
    seenTwice = 0

  nums.forEach((num) => {
    seenOnce = ~seenTwice & (seenOnce ^ num)
    seenTwice = ~seenOnce & (seenTwice ^ num)
  })

  return seenOnce
}
