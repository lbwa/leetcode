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
  - [hard](#hard)

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
- [70 爬楼梯](src/easy-70)
  - 动态规划入门经典
  - 问题核心在于拆分复杂问题为简单问题，到第 n 阶有两种方案：
    1. 先到第 n - 1 阶，再跨 1 阶到第 n 阶
    2. 先到第 n - 2 阶，再跨 2 阶到第 n 阶
    3. 那么合并两种方案就是到第 n 阶的方案数，即 `f(n) = f(n - 1) + f(n - 2)`
    4. 为什么是两种？因为每次只可以走 1 步或 2 步，所以是 n - 1 与 n - 2 之和
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
- [700 二叉搜索树中的搜索](src/easy-700)
  - BST 中，左侧子树节点始终小于当前节点，右侧反之
  - 通过递归实现 BST 搜索
- [703 数据流中的第 K 大的元素](src/easy-703)
  - 通过维护一个 **有限大小，且大小为 k** 的小顶堆，来实现获取数据流中的第 K 大（而不是第 K 个，区别于 [215](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/) 题）元素（含重复元素）。
  - 此时小顶堆为目标值，其所有子节点都比根节点大。

### medium

- [19 删除链表倒数第 k 个节点](src/medium-19)
  - 快慢指针
- [74 搜索二维矩阵](src/medium-74)
  - 基于有序二维矩阵的查找
- [94 二叉树的中序遍历](src/medium-94)
  - 首先递归访问左子树，后访问根节点，最后递归访问右子树
  - 可通过递归或迭代方式实现
- [98 验证二叉搜索树](src/medium-98)
  - 一棵符合定义的二叉搜索树的 **所有** 左子节点小于当前节点，所有右子树节点大于当前节点
  - 由定义，借助递归策略来验证每个节点是否落在特定的取值范围内，即可验证当前二叉树是否为有效的二叉搜索树
    - 所有左子节点都小于父节点，大于 `-Infinity`
    - 所有右子节点都大于父节点，小于 `Infinity`
  - （推荐）直接使用二叉树的中序（迭代）遍历来确定当前节点是否大于前一节点，一旦为 `false`，那么可退出迭代，并返回结果当前树并非二叉搜索树
- [102 二叉树的层序遍历（广度优先搜索 BSF）](src/medium-102)
  - 每次迭代时都取第 `Sn` 层的 `n` 个节点进行迭代，而不再像前中后序遍历那样每次取一个节点。
- [141 环形链表 II](src/medium-142)
  - 找到环形链表的开始节点
- [144 二叉树的前序遍历](src/medium-144)
  - 首先访问根节点，后递归访问左子树，最后递归访问右子树
  - 可通过递归或迭代方式实现
- [230 二叉搜索树中第 K 小的元素](src/medium-230)
  - - [算法 4th][book-algs4] 中常规的 BST `选择 select` 操作
- [347 前 K 个高频元素](src/medium-347)
- [429 n 叉树的层序遍历](src/medium-429)
  - 原理同 `medium 102 二叉树的层序遍历（BSF）`
- [451 根据字符出现的频率排序](src/medium-451)
  - [优先队列][github-priority-queue]（基于二叉堆实现）
- [692 前 K 个高频单词](src/medium-692)
- [701 二叉搜索树中的插入操作](src/medium-701)
  - [算法 4th][book-algs4] 中常规的 BST `插入 put` 操作

### hard

- [145 二叉树的后序遍历](src/hard-145)
  - 首先递归访问左子树，后递归访问右子树，最后访问根节点
  - 可通过递归或迭代方式实现
- [297 二叉树的序列化和反序列化](src/hard-297)
  - 通过前序遍历序列化二叉树，需包含叶子节点的左右 `null` 节点
  - 通过递归反序列化得到二叉树

[github-priority-queue]: https://github.com/lbwa/algorithms/tree/master/data-structures/priority-queue
[book-algs4]: https://algs4.cs.princeton.edu/32bst/
