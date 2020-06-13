import { climbStairs } from '.'

describe('70 爬楼梯', () => {
  it('Should return a correct result', () => {
    expect(climbStairs(1)).toBe(1)
    expect(climbStairs(2)).toBe(2)
    expect(climbStairs(3)).toBe(3)
    expect(climbStairs(4)).toBe(5)
    expect(climbStairs(44)).toBe(1134903170)
    expect(climbStairs(45)).toBe(1836311903)
  })
})
