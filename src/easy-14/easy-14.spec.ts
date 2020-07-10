import { longestCommonPrefix } from '.'

describe('14 最长公共前缀', () => {
  it('Should return correct answers', () => {
    expect(longestCommonPrefix(['flower', 'flow', 'flight'])).toBe('fl')
    expect(longestCommonPrefix(['dog', 'racecar', 'car'])).toBe('')
  })
})
