import { minSubArrayLen } from '.'

describe('209 长度最小的子数组', () => {
  it('Should return a correct result', () => {
    expect(minSubArrayLen(7, [2, 3, 1, 2, 4, 3])).toBe(2)
    expect(minSubArrayLen(100, [])).toBe(0)
    expect(minSubArrayLen(3, [1, 1])).toBe(0)
  })
})
