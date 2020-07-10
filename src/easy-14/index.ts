export function longestCommonPrefix(strs: string[]): string {
  if (strs.length < 1) return ''

  const pivot = strs[0] // 后续项与首项的各个字符对比，即实现纵向扫描
  for (let i = 0; i < strs[0].length; i++) {
    for (let j = 1; j < strs.length; j++) {
      if (i === strs[j].length || strs[j][i] !== pivot[i]) {
        return strs[0].slice(0, i)
      }
    }
  }

  // 迭代完首项的所有字符均相等，即首项为其中最短字符串，那么输入首项作为最长公共项
  return strs[0]
}
