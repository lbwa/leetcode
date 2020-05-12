import { isValid } from './index'

describe('EASY 20. 有效括号', () => {
  ;[
    {
      test: '',
      expected: true
    },
    {
      test: '{}',
      expected: true
    },
    {
      test: '()[]{}',
      expected: true
    },
    {
      test: '(]',
      expected: false
    },
    {
      test: '([)]',
      expected: false
    },
    {
      test: '{[]}',
      expected: true
    }
  ].forEach(({ test, expected }) => {
    it(test, () => {
      expect(isValid(test))[expected ? 'toBeTruthy' : 'toBeFalsy']()
    })
  })
})
