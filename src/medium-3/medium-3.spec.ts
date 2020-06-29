import { lengthOfLongestSubstring } from '.'

describe('3 无重复最长子字符串', () => {
  it('Should return a correct answer', () => {
    expect(lengthOfLongestSubstring('abcabcbb')).toBe(3)
    expect(lengthOfLongestSubstring('bbbbb')).toBe(1)
    expect(lengthOfLongestSubstring('pwwkew')).toBe(3)
  })
})
