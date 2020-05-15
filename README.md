<h1 align="center">leetcode</h1>

<p align="center">
  <a href="https://github.com/lbwa/leetcode/actions">
    <img alt="unit tests" src="https://github.com/lbwa/leetcode/workflows/Unit%20tests/badge.svg">
  </a>
</p>

The solutions for leetcode.com.

## 对撞指针

对撞指针指在有序数组中，将指向最左侧的索引定义 `左指针`，最右侧的索引定义为 `右指针`，然后从两头向中间进行数组遍历。

对撞指针适用于 **有序数组**。也就是说当你遇到题目给定有序数组时，应第一时间想到用对撞指针解题。

```js
function fn(list) {
  let left = 0
  let right = list.length - 1

  while (left <= right) {
    left++
    // do something
    right--
  }
}
```

## 快慢指针

亦称 `双指针`，适用于 **有序数组**。两指针从同侧开始遍历数组。将这两个指针分别定义位 `快指针` 和 `慢指针`。**两个指针以不同的策略移动**，直到两值相等为止。
