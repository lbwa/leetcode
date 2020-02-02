export function reverseWords(str: string) {
  const stack = str.split(' ')
  return stack.reduce((result: string, char: string, index: number) => {
    return (result +=
      (index ? ' ' : '') +
      char
        .split('')
        .reverse()
        .join(''))
  }, '')
}
