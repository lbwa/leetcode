export function rotate(nums: number[], k: number): void {
  let step = k % nums.length
  while (step--) {
    nums.unshift(nums.pop()!)
  }
}
