# 填充每个节点的下一个右侧结点指针 ii

给定一个二叉树

```
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
```

填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有  next 指针都被设置为 NULL。

进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

## References

- [leetcode-cn](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/)
