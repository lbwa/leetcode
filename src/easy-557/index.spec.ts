import { reverseWords } from './index'

describe('给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。', () => {
  it("输入：Let's take LeetCode contest 输出：s'teL ekat edoCteeL tsetnoc", () => {
    const string = "Let's take LeetCode contest"
    const expected = "s'teL ekat edoCteeL tsetnoc"
    expect(reverseWords(string)).toEqual(expected)
  })
})
