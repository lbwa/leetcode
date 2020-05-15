import { sortedSquares } from './index'

describe('easy 977', () => {
  it('Should pass', () => {
    expect(sortedSquares([-4, -1, 0, 3, 10])).toEqual([0, 1, 9, 16, 100])
    expect(sortedSquares([-7, -3, 2, 3, 11])).toEqual([4, 9, 9, 49, 121])
  })
})
