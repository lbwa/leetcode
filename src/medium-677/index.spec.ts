import { MapSum } from '.'

describe('MapSum', () => {
  it('Should return correct answer', () => {
    const obj = new MapSum()
    obj.insert('apple', 3)
    expect(obj.sum('ap')).toBe(3)
    obj.insert('app', 2)
    expect(obj.sum('ap')).toBe(5)
  })
})
