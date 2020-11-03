<h1 align="center">leetcode</h1>

<p align="center">
  <a href="https://github.com/lbwa/leetcode/actions">
    <img alt="unit tests" src="https://github.com/lbwa/leetcode/workflows/Unit%20tests/badge.svg">
  </a>
</p>

`leetcode` 解答。

<!-- TOC -->

- [How to run specific unit-test](#how-to-run-specific-unit-test)
- [树](#%E6%A0%91)
  - [easy](#easy)
  - [medium](#medium)
  - [hard](#hard)
- [链表](#%E9%93%BE%E8%A1%A8)
  - [easy](#easy)
  - [medium](#medium)
  - [hard](#hard)
- [二分查找](#%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE)
  - [easy](#easy)
  - [medium](#medium)
  - [hard](#hard)
- [滑动窗口](#%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3)
  - [medium](#medium)
  - [hard](#hard)
- [双指针](#%E5%8F%8C%E6%8C%87%E9%92%88)
  - [easy](#easy)
  - [medium](#medium)
  - [对撞指针](#%E5%AF%B9%E6%92%9E%E6%8C%87%E9%92%88)
  - [快慢指针](#%E5%BF%AB%E6%85%A2%E6%8C%87%E9%92%88)
- [二叉堆/优先队列](#%E4%BA%8C%E5%8F%89%E5%A0%86%E4%BC%98%E5%85%88%E9%98%9F%E5%88%97)
  - [easy](#easy)
  - [medium](#medium)
- [动态规划](#%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92)
  - [easy](#easy)
  - [medium](#medium)
- [前缀树](#%E5%89%8D%E7%BC%80%E6%A0%91)
  - [medium](#medium)
- [回溯](#%E5%9B%9E%E6%BA%AF)
  - [medium](#medium)
- [散列映射](#%E6%95%A3%E5%88%97%E6%98%A0%E5%B0%84)
  - [easy](#easy)
- [Difficulty level](#difficulty-level)
  - [easy](#easy)

<!-- /TOC -->

## How to run specific unit-test

```bash
$ npm t <DIR_NAME_OR_FILE_NAME>
# eg. to run medium-692
# npm t medium-692
```

[github-priority-queue]: https://github.com/lbwa/algorithms/tree/master/data-structures/priority-queue
[book-algs4-bst]: https://algs4.cs.princeton.edu/32bst/
[book-algs4-pq]: https://algs4.cs.princeton.edu/24pq/

## 树

> tree traversal

### easy

- [26 删除有序数组中的重复项](src/easy-26)
  - 在处理数组项删除问题时，尽量避免直接删除数组中间项，而是排列到尾部删除
  - 将所有重复项向尾部排列
  - 维护 `[0, ... slow]` 即可实现重复项删除
- [83 删除有序链表中的重复元素](src/easy-83)
  - 原理同 `26 删除有序数组中的重复项`
  - 维持 `头部 ~ slow` 部分为唯一项部分
- [100 相同的树](src/easy-100)
  - 根据递归法遍历树节点逐一验证每个节点是否相等
- [101 对称二叉树](src/easy-101)
- [104 二叉树的最大深度](src/easy-104)
  - 法一：通过 DFS 遍历节点得到最大深度
  - 法二：通过 BFS 遍历节点得到最大深度
- [108 将有序数组转换为二叉搜索树](src/easy-108)
  - 要构建一个平衡二叉搜索树的关键在于使用 `BST` 且使用有序数组的中项作为根节点
- [110 平衡二叉树](src/easy-110)
  - 本质是通过 `DFS` 来确定左右子树的高度（最大深度）差值不大于 1。
- [111 二叉树的最小深度](src/easy-111)
  - 递归遍历二叉树，得到最小深度
  - 迭代法：借助 `429 二叉树层序遍历` 的框架实现探测最小深度，当某层某节点为叶子元素时，返回当前深度
- [226 翻转二叉树](src/easy-226)
  - 递归（迭代）遍历二叉树度
- [236 二叉搜索树的最近公共祖先](src/easy-236)
  - 据题，给定节点一定在树中；
  - 利用 BST 的特点，缩小查找范围
  - 当遍历得到一个节点的值在给定两节点之间时，在 BST 这样的结构中一定是给定两节点的最近公共祖先
- [559 n 叉树的最大深度](src/easy-559)
  - 递归（迭代）遍历 n 叉树
- [700 二叉搜索树中的搜索](src/easy-700)
  - BST 中，左侧子树节点始终小于当前节点，右侧反之
  - 通过递归实现 BST 搜索
- [112 路径总和](src/easy-112)
  - 层序遍历实现路径求和
- [257 二叉树的所有路径](src/easy-257)
  - 即仅在遍历至叶子节点时，`push` 路径，否则继续加长路径
- [404 左叶子之和](src/easy-404)
  - 注意本题为所有 **左叶子** 节点而非所有左节点之和
  - 常规树的递归遍历得到左叶子之和
- [530 二叉搜索树的最小绝对差](src/easy-530)
  - 因为是二叉搜索树，所以在中序遍历的情况下，遍历值为递增序列，故可据此得到最小差
  - 基于中序遍历(DFS) 实现两两节点求差，得到最小差
- [538 把二叉搜索树转换为累加树](src/easy-538)
  - 同 `1038`，反向中序遍历得到结果树
- [617 合并二叉树](src/easy-617)
  - 递归前序遍历合并二叉树
- [637 二叉树层的平均值](src/easy-637)
  - 通过二叉树的层序遍历得到每层节点和与节点数，在求平均值前，将前二者转换为 `float64` 类型后，再求得当前层的平均值
- [783 二叉搜索树节点的最小距离](src/easy-783)
  - 同 `530 二叉搜索树的最小绝对差`
  - 因为是二叉搜索树，所以在中序遍历的情况下，遍历值为递增序列，故可据此得到最小差
  - 基于中序遍历(DFS) 实现两两节点求差，得到最小差

### medium

- [94 二叉树的中序遍历](src/medium-94)
  - 首先递归访问左子树，后访问根节点，最后递归访问右子树
  - 可通过递归或迭代方式实现
- [98 验证二叉搜索树](src/medium-98)
  - 一棵符合定义的二叉搜索树的 **所有** 左子节点小于当前节点，所有右子树节点大于当前节点
  - 由定义，借助递归策略来验证每个节点是否落在特定的取值范围内，即可验证当前二叉树是否为有效的二叉搜索树
    - 所有左子节点都小于父节点，大于 `-Infinity`
    - 所有右子节点都大于父节点，小于 `Infinity`
  - （推荐）直接使用二叉树的中序（迭代）遍历来确定当前节点是否大于前一节点，一旦为 `false`，那么可退出迭代，并返回结果当前树并非二叉搜索树
- [102 二叉树的层序遍历（广度优先搜索 BFS）](src/medium-102)
  - 每次迭代时都取第 `Sn` 层的 `n` 个节点进行迭代，而不再像前中后序遍历那样每次取一个节点。
- [103 二叉树的锯齿形层次遍历](src/medium-103)
  - 按照标准 `BFS` 遍历后，对奇数层结果进行翻转。
- [105 从前序与中序遍历序列构造二叉树](src/medium-105)
  - 根据前序结果的首项得到第一个根节点，根据根节点，分别切分左右结果数组，递归生成左右子树的根节点，最终得到结果二叉树
- [106 从中序与后序遍历序列构造二叉树](src/medium-106)
  - 根据后序结果数组的末项得到第一个根节点，根据根节点，分别切分左右结果数组，递归生成左右子树的根节点，最终得到结果二叉树
- [109 将有序链表转换二叉搜索树](src/medium-109)
  - 原理同 `easy 104`，为了构建平衡的二叉搜索树，那么取链表的中项（通过快慢指针找到）作为左右子树的根节点。
- [113 路径总和 ii](src/medium-113)
  - 前序 DFS，并在每次遍历当前节点时判断是否符合要求，否则继续遍历所有子树，并在之后 `pop` 当前节点，以避免在后续路径中出现该节点
- [114 二叉树展开为链表](src/medium-114)
  - 基于二叉树前序遍历，每次遍历时，缓存当前节点，以用于将当前节点重新连接下一节点。
- [116 填充每个节点的下一个右侧节点指针](src/medium-116)
  - 基于层序遍历迭代法实现
- [117 填充每个节点的下一个右侧节点指针 II](src/medium-116)
  - 基于层序遍历迭代法实现
- [129 求根到叶子节点数字之和](src/medium-129)
  - 同 `medium 113 路径总和 ii` 的 `DFS` 核心思路，最终根据得到的路径相加得到结果。
- [144 二叉树的前序遍历](src/medium-144)
  - 首先访问根节点，后递归访问左子树，最后递归访问右子树
  - 可通过递归或迭代方式实现
- [199 二叉树的右视图](src/medium-199)
  - 基于层序遍历迭代法实现，仅保存当层节点的最后一个节点，最终所有层的最后一个节点构成二叉树的右视图
- [230 二叉搜索树中第 K 小的元素](src/medium-230)
  - [算法 4th - BST][book-algs4-bst] 中常规的 BST `选择 select` 操作
- [236 二叉树的最近公共祖先](src/medium-236)
  - 基于（递归）后序遍历，那么在所有符合条件的公共祖先中，一定是最近的公共祖先首先被访问到。
  - 在递归结构中，始终传入 `p`, `q` 节点，并与之与当前节点的左右子树中进行查找，当左子树包含其中一个节点时，那么右子树至多包含另一节点。
- [429 n 叉树的层序遍历](src/medium-429)
  - 原理同 `medium 102 二叉树的层序遍历（BFS）`
- [662 二叉树最大宽度](src/medium-662)
  - 借鉴在 `PriorityQueue` 及其相似二叉树结构通过顺序存储时，`index` 节点的左子节点为 `2*index`, 右子节点为 `2*index + 1`<sup>[算法 4th - PQ][book-algs4-pq]</sup>
  - 通过层序遍历（广度优先搜索，BFS）时队列中的首项和末项的索引差值加一得到二叉树的最大宽度（最底层叶子节点的跨度）。
- [701 二叉搜索树中的插入操作](src/medium-701)
  - [算法 4th - BST][book-algs4-bst] 中常规的 BST `插入 put` 操作
- [889 根据前序和后序遍历构造二叉树](src/medium-889)
  - 根据前序遍历的后一节点得到左子树的根节点的值，据此值，切分后序遍历结果数组，递归生成结果子树
- [1008 先序遍历构造二叉搜索树](src/medium-1008)
  - 本质上是 [算法 4th - BST][book-algs4-bst] 中常规的 BST 的 `插入 put` 节点操作
- [1038 从二叉搜索树到更大和树](src/medium-1038)
  - 同 `538`，反向中序遍历得到结果树
- [1609 奇偶数](src/medium-1609)
  - 基于层序遍历（BST）逐层检测节点是否符合奇偶数定义

### hard

- [145 二叉树的后序遍历](src/hard-145)
  - 首先递归访问左子树，后递归访问右子树，最后访问根节点
  - 可通过递归或迭代方式实现
- [297 二叉树的序列化和反序列化](src/hard-297)
  - 通过前序遍历序列化二叉树，需包含叶子节点的左右 `null` 节点
  - 通过递归反序列化得到二叉树

## 链表

> linked list

### easy

- [83 删除有序链表中的重复元素](src/easy-83)
  - 原理同 `26 删除有序数组中的重复项`
  - 维持 `头部 ~ slow` 部分为唯一项部分
- [141 环形链表](src/easy-141)
  - 快慢指针
- [203 移除链表元素](src/easy-203)
  - 使用哨兵节点（策略）实现规避头节点的复杂情况处理
  - 使用快慢指针实现链表节点删除
- [206 反转单向链表](src/easy-206)
  - 迭代法，基于常规的链表迭代：
    1. 首先保存下一节点的链接；
    1. 反转当前节点的下一链接至上一节点；
    1. 将 prev 保存为上一节点；
    1. 继续迭代 1 中保存的原有链表的下一节点
  - 递归法，基于递归栈，在栈帧中，将下一节点的下一节点指向当前节点，即实现反转，另需要断开原有链接的 `next` 链接
- [234 回文链表](src/easy-234)
  - 时间复杂度为 `O(N)`，空间复杂度为 `O(1)` 的解法：
    1. 快慢双指针找到链表中项
    1. 反转中项（不含中项）至尾项部分链表
    1. **以后半部分链表为基准**，双指针迭代比较 `i`（`[0, mid]`） 和 `j`（`(mid, end]`） 项的值是否满足 "回文" 定义
       > 因为以后半部分链表为基准，那么当项数为奇数时不会迭代中项，当项数为偶数时，会迭代中项。
  - 时间，空间复杂度均为 `O(N)` 的解法：
    - 基于存在二叉树的特例即为单向列表，那么借助 `二叉树的后续遍历` 和 `双指针` 策略可实现回文链表检测

### medium

- [2 两数相加](src/medium-2)
  - 题为低位在链表头，那么两指针同时移动至链表尾，过程中，逐位相加，且行走至链表尾的指针不再移动，相加时补零
  - 通过尾插法（`medium 445 两数之和II` 为头插法）生成结果链表
- [19 删除链表倒数第 k 个节点](src/medium-19)
  - 哨兵节点开始，快指针先走 `k+1` 步后，快慢指针齐走，至快指针为 `nil` 时，慢指针的下一节点即为待删除节点
- [24 两两交换链表中的节点](src/medium-24)
  - 基于头插法两两反转
  - 每隔一个节点开始两两反转
- [74 搜索二维矩阵](src/medium-74)
  - 基于有序二维矩阵的查找
- [82 删除有序链表中的重复项 II](src/medium-82)
  - 哨兵 `sentinel` 节点用于处理链表节点删除的场景题，规避头节点的复杂情况
  - 双指针，`current` 节点始终为非重复项节点，迭代比较 `next` 和 `next.next` 节点：
    - 若发生重复，那么通过临时指针 `temp` 找到下一个没有重复项的节点或 `null` 作为
      `current` 的 `next` 节点
    - 若没有重复，那么迭代 `current` 的 `next` 节点
  - 结果为 `sentinel` 节点的 `next` 节点为头节点的链表
- [92 反转链表 II](src/medium-92)，即反转 `[m, n]` 链表项
  - 递归法，基于反转前 `n` 项递归法，而此基于反转链表递归法
  - 迭代法：
    1. 先走 `m - 1` 步；
    1. 反转单项；
    1. 保存待反转项；
    1. 将待反转项从链表中分离
    1. 反转待反转项至前项
    1. 将反转后的项重新加入链表中
- [142 环形链表 II](src/medium-142)
  - 找到环形链表的开始节点
  - 基于 `141 环形链表` 解法判断有环
  - 重置其中一个指针至头节点，同步走两个指针，此时双指针碰撞时，即为环的开始节点
- [445 两数相加 II](src/medium-445)
  - 高位在链表表头，故首先创建两个栈维护链表节点，借鉴 `medium 2 两数相加` 的思路逐位相加，且通过头插法生成结果链表。
- [707 设计链表](src/medium-707)
  - 在设计之初，头节点应始终为一个哨兵节点，以排除后续对头节点的操作

### hard

- [25 K 个一组翻转链表](src/hard-25)
  - `24 两两交换链表中的节点` 的通用版本，`92 反转链表 ii` 的多次操作版本
  - 可看作执行了多次链表区间反转
  - 在每次执行 `k` 个反转前，验证是否存在 `k` 个节点，不足时直接返回，不再反转不够 `k` 个节点的区间

## 二分查找

### easy

- [35 搜索插入位置](src/easy-35)
  - 经典二分查找
- [704 二分查找](src/easy-704)
  - 经典二分查找

### medium

- [33 搜索有序数组](src/medium-33)
  - 首先判断在前半段还是后半段递增序列中，然后再通过朴素二分查找在递增序列中查找
- [153 寻找旋转排序数组中的最小值（不含重复元素）](src/medium-153)
  - 同 `154`

### hard

- [154 寻找旋转排序数组种的最小值（含重复元素）](src/hard-154)
  - 同 `153`
  - 与末项相比，当大于末项时移动左指针，小于末项时移动右指针，否则逐步移动右指针
  - 指针碰撞时，即为最小项

## 滑动窗口

> sliding window

### medium

- [3 无重复字符的最长子串](src/medium-3)
  - 构建字母频率哈希映射
  - 当存在频率大于 1 的字母时，缩小窗口，直到当前字母的频率等于 1
- [209 长度最小的子数组](src/medium-209)
  - 在增大窗口时，增大和；当和满足条件时，缩小窗口
- [438 找到字符串中所有字母异位词](src/medium-438)
  - 基于 `easy 242` 找到指定单词异位词的方法 —— 哈希映射
  - 在窗口大小不小于指定字符串的长度时，那么可能存在异位词，那么开始尝试缩小窗口，并比较 `validsInWindow` 是否等于目标哈希映射的大小，反之，不存在异位词，继续滑动窗口。
- [567 字符串的排列](src/medium-567)
  - `s2` 是否包含 `s1` 的排列，本质上，是求 `s1` 是否是 `s2` 的某一子串，即是求是否存在一个 `s2` 窗口是 `s1` 的异位词，而异位词的求解（`easy 242`）是基于哈希表的。

### hard

- [76 最小覆盖子串](src/hard-76)
  - 利用哈希表记录字符频率来判断子串
  - 通过计数有效字符频率相等时的字符个数（即此时完全覆盖了子串）来缩小窗口
  - 在缩小窗口时应减去计数

## 双指针

> two pointers

### easy

- [9 回文数](src/easy-9)
- [167 两数之和 II - 输入有序数组](src/easy-167)
  - 基于双指针或二分查找的解法，其中只需要返回第一次结果
- [344 反转字符串](src/easy-344)
  - 基于双指针实现高低位索引交换，进而实现反转字符串
- [349 两个数组的交集](src/easy-349)
  - 基于有序数组和 `Set` 的双指针解法，其中结果不含重复项
- [350 两个数组的交集 II](src/easy-350)
  - 基于有序数组的双指针解法，其中结果含重复项
- [925 长按键入](src/easy-925)
  - 基于双指针，分别在目标字符串和实际输入字符串上移动

### medium

- [11 盛水最多的容器](src/medium-11)
  - 对撞双指针，找到矩形面积最大区间
  - 每次移动 `y` 轴值最小的指针
- [15 三数之和](src/medium-15)
  - 通过迭代跳过重复项
  - 将三数之和问题转换为一个数和两数指和的和的问题
  - 两数之和通过双指针求得所有符合为 `0 - 第一个数` 的条件的集合
  - 举一反三地，可将 `n 数之和` 转换为 `1 + n-1 数之和`。
- [16 最接近的三数之和](src/medium-16)
  - 思路与三数之和雷同，基于有序数组的双指针解法
- [18 四数之和](src/medium-15)
  - 在 `medium 15 三数之和` 的基础上新增一个循环
- [75 颜色分类](src/medium-75)
  - 要解决的本质问题是，在不使用排序函数的前提下，原地排序原始数组，使得相同类的颜色相邻。

双指针策略包含快慢指针和对撞指针（亦称左右指针）策略。

### 对撞指针

对撞指针指在有序数组中，将指向最左侧的索引定义 `左指针`，最右侧的索引定义为 `右指针`，然后从两头向中间进行数组遍历。

对撞指针适用于 **有序的数据结构**（如链表，数组等）。

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

亦称 `双指针`，适用于 **有序的数据结构**（如链表，数组等）。两指针从同侧开始遍历数组。将这两个指针分别定义位 `快指针` 和 `慢指针`。**两个指针以不同的策略移动**，直到两值相等为止。

## 二叉堆/优先队列

> binary heap/priority queue

### easy

- [703 数据流中的第 K 大的元素](src/easy-703)
  - 通过维护一个 **有限大小，且大小为 k** 的小顶堆，来实现获取数据流中的第 K 大（而不是第 K 个，区别于 [215](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/) 题）元素（含重复元素）。
  - 此时小顶堆为目标值，其所有子节点都比根节点大。

### medium

- [215 数组中的第 K 个最大元素](src/medium-215)
  - 通过快排或维护一个 `K` 大的二叉堆实现，找到第 K 个最大元素
- [347 前 K 个高频元素](src/medium-347)
- [451 根据字符出现的频率排序](src/medium-451)
  - [优先队列][github-priority-queue]（基于二叉堆实现）
- [692 前 K 个高频单词](src/medium-692)

## 动态规划

> dynamic programming(DP)

解决动态规划问题的核心在于找到对应的 **动态方程**。

### easy

- [70 爬楼梯](src/easy-70)
  - 动态规划入门经典
  - 问题核心在于拆分复杂问题为简单问题，到第 n 阶有两种方案：
    1. 先到第 n - 1 阶，再跨 1 阶到第 n 阶
    2. 先到第 n - 2 阶，再跨 2 阶到第 n 阶
    3. 那么合并两种方案就是到第 n 阶的方案数，即 `f(n) = f(n - 1) + f(n - 2)`
    4. 为什么是两种？因为每次只可以走 1 步或 2 步，所以是 n - 1 与 n - 2 之和
- [118 杨辉三角](src/easy-118)
  - 第 i 行有 i+1 项
  - 边界项始终为 1，即 `dp[i][0] = 1, dp[i][i] = 1`
  - `j > 0 && j < i, dp[i][j] = dp[i-1][j-1] + dp[i-1][j]`
- [119 杨辉三角](src/easy-119)
  - 基于 `118` 的 `DP` 思路和滚动数组实现 `O(k)` 的空间复杂度优化。

### medium

- [62 不同路径](src/medium-62)
  - 典型动态规划问题，每次只移动一步时，得到动态方程 `f(m, n) = f(m - 1, n) + f(m, n - 1)`，即从 `(0, 0)` 到 `(m, n)` 点的所有走法等于从原点到 `(m - 1, n)` 和到 `(m, n - 1)` 之和。
  - 通过 `滚动数组` 优化空间复杂度，实现在迭代过程中覆写中间结果，达到优化空间的目的。
- [63 不同路径 II（含障碍点）](src/medium-63)
  - `62` 的升级版，含障碍点
  - 与 `62` 的核心解法一致，推导得出状态方程，在实现状态方程的基础上，将障碍点的方法步数设为常量 `0`。
  - 通过滚动数组优化空间复杂度，实现在迭代过程中覆写中间结果，达到优化空间的目的。
- [64 最小路径和](src/medium-64)
  - 与 `62` 的状态方程类似，不同的是，矩阵项的值为 `grid` 参数矩阵中路径上的项的值的和。

## 前缀树

> trie tree，亦称字典树

### medium

- [208 实现 Trie 前缀树](src/medium-208)
  - 实现前缀树的关键在于迭代创建字典前缀 map
- [677 键值映射](src/medium-677)
  - 通过 `trie tree` 的 `insert` 实现节点新建
  - 通过 `trie tree` 的 `search` 实现找到以参数变量 `prefix` 开头的前缀节点，通过 `DFS` 求得该节点的所有后代节点的 `val` 之和，最终得到 `sum` 方法的输出

## 回溯

> backtrack

### medium

- [39 组合总和](src/medium-39)
  - 典型组合型回溯算法
  - 枚举过程本身可看作是决策树遍历
- [46](src/medium-46)
  - 在路径项等于原始序列长度时，即表示已经存在一个排列，此时将此路径经加入结果中
  - 在递归过程中，分为选择当前项和不选择当前项（回溯）两条路径
- [78 子集](src/medium-78)
  - 典型组合型回溯算法

## 散列映射

> hash mapping，散列映射，哈希

### easy

- [771 宝石与石头](src/easy-771)
  - 使用哈希映射记录宝石种类，迭代石头得到最终的宝石数量

## Difficulty level

### easy

- [1 两数之和](src/easy-1)
- [14 最长公共前缀](src/easy-14)
  - 以首项作为基准项，通过纵向扫描实现找到最长公共前缀
- [20 有效括号](src/easy-20)
- [242 有效的字母异位词](src/easy-242)
  - 通过 26 个字母哈希表实现字母频率归零，反之为非异位词
- [415 字符串相加](src/easy-415)
  - 模拟十进制数字相加，不够位则自动补零
- [557 反转字符串中的单词 III](src/easy-557)
- [657 机器人能否返回原点](src/easy-657)
  - 关键在于量化机器人的移动步骤，左右，上下相加时可相互抵消，最终验证水平和垂直方向上的和是否为 0
- [941 有效的山脉数组](src/easy-941)
  - 运用一个指针根据题意进行线性扫描得到结果
