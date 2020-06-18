import { removeDuplicates } from '.'

describe('26 删除排序数组中的重复项', () => {
  it('Should return a correct list', () => {
    expect(removeDuplicates([])).toBe(0)
    expect(removeDuplicates([1])).toBe(1)
    const list = [1, 1, 2, 2]
    const answer = removeDuplicates(list)
    expect(answer).toEqual(2)
    expect(list).toEqual([1, 2])
  })
})
