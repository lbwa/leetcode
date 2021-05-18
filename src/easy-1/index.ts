export function twoSum(nums: number[], target: number) {
  const map: Record<string, number> = {}
  let i = nums.length
  while (i--) {
    const current = nums[i]
    if (map[target - current] !== undefined) {
      return [i, map[target - current]]
    }
    map[current] = i
  }
  return [-1, -1]
}

// export function twoSum(nums: number[], target: number) {
//   const map: Record<string, number> = {}
//   return nums.reduce<number[]>((result, cur, index) => {
//     if (map[target - cur] !== undefined) {
//       return [map[target - cur], index]
//     } else {
//       map[cur] = index
//     }
//     return result
//   }, [])
// }
