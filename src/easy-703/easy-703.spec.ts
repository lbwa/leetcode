import { KthLargest } from '.'

describe('703 数据流中的第 K 大的元素', () => {
  it('Should return', () => {
    const handler = new KthLargest(3, [4, 5, 8, 2])
    expect(handler.add(3)).toBe(4)
    expect(handler.add(5)).toBe(5)
    expect(handler.add(10)).toBe(5)
    expect(handler.add(9)).toBe(8)
    expect(handler.add(4)).toBe(8)
  })
})
