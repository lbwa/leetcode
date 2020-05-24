import { topKFrequent } from './index'

describe('347. 前 K 个高频元素', () => {
  it('Should return [4,-1]', () => {
    const list = [2, 3, 4, 1, 4, 0, 4, -1, -2, -1]
    expect(topKFrequent(list, 2)).toEqual([4, -1])
  })

  it('Should return [1, 2]', () => {
    const list = [1, 1, 1, 2, 2, 3]
    expect(topKFrequent(list, 2)).toEqual([1, 2])
  })
})
