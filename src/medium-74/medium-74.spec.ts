import { searchMatrix } from './index'

describe('74 搜索二维矩阵', () => {
  it('[]', () => {
    expect(searchMatrix([], 1)).toBeFalsy()
    expect(searchMatrix([[]], 1)).toBeFalsy()
  })

  it('[[1,3,5,7], [10,11,16,20], [23,30,34,50]] 5', () => {
    expect(
      searchMatrix(
        [
          [1, 3, 5, 7],
          [10, 11, 16, 20],
          [23, 30, 34, 50]
        ],
        3
      )
    ).toBeTruthy()
  })

  it('[[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 50]] 13', () => {
    expect(
      searchMatrix(
        [
          [1, 3, 5, 7],
          [10, 11, 16, 20],
          [23, 30, 34, 50]
        ],
        13
      )
    ).toBeFalsy()
  })
})
