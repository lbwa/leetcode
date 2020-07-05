import { uniquePathsSpaceO2n, uniquePathsSpaceOn } from '.'

describe('62 不同路径', () => {
  it('Should return correct answers', () => {
    expect(uniquePathsSpaceO2n(3, 2)).toBe(3)
    expect(uniquePathsSpaceO2n(7, 3)).toBe(28)

    expect(uniquePathsSpaceOn(3, 2)).toBe(3)
    expect(uniquePathsSpaceOn(7, 3)).toBe(28)
  })
})
