import { frequencySort } from './index'

describe('451. 根据字符出现频率排序', () => {
  it('Should return eert', () => {
    expect(frequencySort('tree')).toEqual('eert')
  })

  it('Should return aaaccc', () => {
    expect(frequencySort('cccaaa')).toEqual('aaaccc')
    expect(frequencySort('aaaccc')).toEqual('aaaccc')
  })

  it('Should return bbAa', () => {
    expect(frequencySort('Aabb')).toEqual('bbAa')
  })
})
