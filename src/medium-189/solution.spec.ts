import { rotate } from './solution'

describe('medium 189', () => {
  it('should be passed', () => {
    const a0 = [1, 2]
    const a1 = [1, 2, 3, 4, 5, 6, 7]
    rotate(a0, 3)
    rotate(a1, 3)
    expect(a0).toEqual([2, 1])
    expect(a1).toEqual([5, 6, 7, 1, 2, 3, 4])
  })
})
