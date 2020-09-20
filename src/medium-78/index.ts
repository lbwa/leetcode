export function subsets0(nums: number[]): number[][] {
  const answer = [] as number[][]
  const n = nums.length
  for (let mask = 0; mask < 1 << n; mask++) {
    const set = [] as number[]
    for (let i = 0; i < n; i++) {
      if ((mask >> i) & 1) {
        set.push(nums[i])
      }
    }
    answer.push(set)
  }
  return answer
}
