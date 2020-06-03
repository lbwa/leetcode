# 703 数据流中的第 K 大的元素

设计一个找到数据流中第 K 大元素的类（class）。**注意是排序后的第 K 大元素，不是第 K 个不同的元素**，也不是如 [215](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/) 一样的第 k 个最大元素。

你的  KthLargest  类需要一个同时接收整数  k 和整数数组 nums  的构造器，它包含数据流中的初始元素。每次调用  KthLargest.add，返回当前数据流中第 K 大的元素。

示例:

```java
int k = 3;
int[] arr = [4,5,8,2];
KthLargest kthLargest = new KthLargest(3, arr);
kthLargest.add(3);   // returns 4
kthLargest.add(5);   // returns 5
kthLargest.add(10);  // returns 5
kthLargest.add(9);   // returns 8
kthLargest.add(4);   // returns 8
```

说明:
你可以假设  nums  的长度 ≥ k-1  且 k ≥ 1。

## References

- [leetcode-cn](https://leetcode-cn.com/problems/kth-largest-element-in-a-stream)
