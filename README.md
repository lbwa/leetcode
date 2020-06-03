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
- [101 对称二叉树](src/easy-101)
- [104 二叉树的最大深度](src/easy-104)
  - 递归（迭代）遍历二叉树
- [111 二叉树的最小深度](src/easy-111)
  - 递归（迭代）遍历二叉树
- [141 环形链表](src/easy-141)
  - 快慢指针
- [203 移除链表元素](src/easy-203)
  - 使用哨兵节点（策略）实现规避头节点的复杂情况处理
  - 使用快慢指针实现链表节点删除
- [206 反转单向链表](src/easy-206)
- [226 翻转二叉树](src/easy-226)
  - 递归（迭代）遍历二叉树
- [557 反转字符串中的单词 III](src/easy-557)
- [559 n 叉树的最大深度](src/easy-559)
  - 递归（迭代）遍历 n 叉树
- [703 数据流中的第 K 大的元素](src/easy-703)
  - 通过维护一个 **有限大小，且大小为 k** 的小顶堆，来实现获取数据流中的第 K 大（而不是第 K 个，区别于 [215](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/) 题）元素（含重复元素）。
  - 此时小顶堆为目标值，其所有子节点都比根节点大。

### medium

- [19 删除链表倒数第 k 个节点](src/medium-19)
  - 快慢指针
- [74 搜索二维矩阵](src/medium-74)
  - 基于有序二维矩阵的查找
- [141 环形链表 II](src/medium-142)
  - 找到环形链表的开始节点
- [347 前 K 个高频元素](src/medium-347)
- [451 根据字符出现的频率排序](src/medium-451)
  - [优先队列][github-priority-queue]（基于二叉堆实现）
- [692 前 K 个高频单词](src/medium-692)

[github-priority-queue]: https://github.com/lbwa/algorithms/tree/master/data-structures/priority-queue
