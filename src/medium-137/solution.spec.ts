import { singleNumber } from './solution'

describe('medium 137 single number ii', () => {
  it('should be passed', () => {
    expect(singleNumber([2, 2, 3, 2])).toEqual(3)
    expect(singleNumber([0, 1, 0, 1, 0, 1, 99])).toEqual(99)
  })
})
