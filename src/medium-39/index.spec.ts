import { combinationSum } from '.'

describe('Test combinationSum', () => {
  it('Should return answer', () => {
    expect(combinationSum([2, 3, 6, 7], 7)).toStrictEqual([[7], [2, 2, 3]])
    expect(combinationSum([2, 3, 5], 8)).toStrictEqual([
      [3, 5],
      [2, 3, 3],
      [2, 2, 2, 2]
    ])
  })
})
