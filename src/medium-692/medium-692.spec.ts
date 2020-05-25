import { topKFrequent } from './index'

describe('692. 前K个高频单词 ', () => {
  it('Should return  [ "i", "love", "coding" ]', () => {
    expect(
      topKFrequent(['i', 'love', 'leetcode', 'i', 'love', 'coding'], 3)
    ).toEqual(['i', 'love', 'coding'])
  })

  it('Should return ["a", "aa"]', () => {
    expect(topKFrequent(['a', 'aa', 'aaa'], 2)).toEqual(['a', 'aa'])
  })

  it('Should return ["a"]', () => {
    expect(topKFrequent(['aaa', 'aa', 'a'], 1)).toEqual(['a'])
  })
})
