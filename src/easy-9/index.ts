export function isPalindrome(number: number) {
  if (number < 0) return false

  const string = number + ''

  for (let i = 0, len = string.length; i < len; i++) {
    const fromHead = string[i]
    const fromTail = string[len - 1 - i]
    if (fromHead !== fromTail) {
      return false
    }
  }

  return true
}

// var isPalindrome = function(x) {
//   return x == x.toString().split('').reverse().join('');
//  };
// 思路一：
// 1.1.回文即正反读值不变，则考虑将数值转化为数组使用reverse()方法在合并达到反转数值的目的
// 1.2.比较原值与翻转后的值
// 思路二：
// 回文第一项字符与最后一项字符相等
