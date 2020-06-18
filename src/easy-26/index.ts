export function removeDuplicates(nums: number[]): number {
  if (nums.length < 1) return nums.length

  let slow = 0,
    fast = 1
  const len = nums.length
  while (fast < len) {
    if (nums[slow] !== nums[fast]) {
      slow++
      nums[slow] = nums[fast]
    }

    fast++
  }

  return (nums.length = slow + 1)
}
