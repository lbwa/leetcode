<h1 align="center">leetcode</h1>

<p align="center">
  <a href="https://github.com/lbwa/leetcode/actions">
    <img alt="unit tests" src="https://github.com/lbwa/leetcode/workflows/Unit%20tests/badge.svg">
  </a>
</p>

The solutions for leetcode.com.

<!-- TOC -->

- [How to run specific unit-test](#how-to-run-specific-unit-test)
- [双指针策略](#%E5%8F%8C%E6%8C%87%E9%92%88%E7%AD%96%E7%95%A5)
  - [对撞指针](#%E5%AF%B9%E6%92%9E%E6%8C%87%E9%92%88)
  - [快慢指针](#%E5%BF%AB%E6%85%A2%E6%8C%87%E9%92%88)
- [Levels](#levels)
  - [easy](#easy)
  - [medium](#medium)

<!-- /TOC -->

## How to run specific unit-test

```bash
$ npm t <DIR_NAME_OR_FILE_NAME>
# eg. to run medium-692
# npm t medium-692
```

## 双指针策略

双指针策略包含快慢指针和对撞指针（亦称左右指针）策略。

### 对撞指针

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

### 快慢指针

亦称 `双指针`，适用于 **有序数组**。两指针从同侧开始遍历数组。将这两个指针分别定义位 `快指针` 和 `慢指针`。**两个指针以不同的策略移动**，直到两值相等为止。

## Levels

### easy

- [1 两数之和](src/easy-1)
- [9 回文数](src/easy-9)
- [20 有效括号](src/easy-20)
- [104 二叉树的最大深度](src/easy-104)
- [206 反转单向链表](src/easy-206)
- [557 反转字符串中的单词 III](src/easy-557)
- [559 n 叉树的最大深度](src/easy-559)

### medium

- [19 删除链表倒数第 k 个节点](src/medium-19)
- [347 前 K 个高频元素](src/medium-347)
- [451 根据字符出现的频率排序](src/medium-451)
- [692 前 K 个高频单词](src/medium-692)
