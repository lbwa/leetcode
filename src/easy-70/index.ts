/**
 * @core 爬楼梯第 n 阶的方案可拆分两种大类方案，如下：
 * 1. 到第 n - 1 阶，然后再跨 1 阶到第 n 阶
 * 2. 到第 n - 2 阶，然后再跨 2 阶到第 n 阶
 * 那么按照常规逻辑，合并以上两种大类方案即可得到 n 阶楼梯的方案数，而第 1 类对应方案
 * 数为 f(n - 1)，第 2 类方案数为 f(n - 2)，所以有 f(n) = f(n - 1) + f(n - 2)
 * @param n 目标台阶数
 */
export function climbStairs(n: number): number {
  let i = 0, // 第 n - 2 个台阶的方案数
    j = 0, // 第 n - 1个台阶的方案数
    m = 1 // 第 n 个台阶的方案数

  for (let stairs = 0; stairs < n; stairs++) {
    i = j
    j = m
    m = i + j
  }

  return m
}
