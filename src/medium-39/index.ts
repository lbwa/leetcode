export function combinationSum(
  candidates: number[],
  target: number
): number[][] {
  const answer: number[][] = []
  function dfs(target: number, index: number, combine: number[]) {
    if (index === candidates.length) {
      return
    }
    if (target === 0) {
      answer.push(combine)
      return
    }
    dfs(target, index + 1, combine)
    if (target - candidates[index] >= 0) {
      dfs(target - candidates[index], index, [...combine, candidates[index]])
    }
  }
  dfs(target, 0, [])
  return answer
}
