import {
  uniquePathsWithObstaclesSpaceO2n,
  uniquePathsWithObstaclesSpaceOn
} from '.'

describe('63 不同路径', () => {
  it('时间复杂度 O(cols * rows)，空间复杂度 O(2n)', () => {
    expect(uniquePathsWithObstaclesSpaceO2n([[1]])).toBe(0)
    expect(uniquePathsWithObstaclesSpaceO2n([[1, 0]])).toBe(0)
    expect(
      uniquePathsWithObstaclesSpaceO2n([
        [0, 0, 0],
        [0, 1, 0],
        [0, 0, 0]
      ])
    ).toBe(2)
    expect(
      uniquePathsWithObstaclesSpaceO2n([
        [0, 0],
        [1, 0]
      ])
    ).toBe(1)
  })

  it('时间复杂度 O(cols * rows)，空间复杂度 O(n)', () => {
    expect(uniquePathsWithObstaclesSpaceOn([[1]])).toBe(0)
    expect(uniquePathsWithObstaclesSpaceOn([[1, 0]])).toBe(0)
    expect(
      uniquePathsWithObstaclesSpaceOn([
        [0, 0, 0],
        [0, 1, 0],
        [0, 0, 0]
      ])
    ).toBe(2)
    expect(
      uniquePathsWithObstaclesSpaceOn([
        [0, 0],
        [1, 0]
      ])
    ).toBe(1)
  })
})
