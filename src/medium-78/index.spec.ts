import { subsets0 } from '.'

describe('subsets', () => {
  it('Should return a correct result', () => {
    expect(subsets0([1, 2, 3])).toStrictEqual([
      [],
      [1],
      [2],
      [1, 2],
      [3],
      [1, 3],
      [2, 3],
      [1, 2, 3]
    ])
  })
})
