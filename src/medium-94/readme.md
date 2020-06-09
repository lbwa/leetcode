# 94 二叉树的中序遍历

给定一个二叉树，返回它的中序遍历。

示例:

输入: [1,null,2,3]

```
   1
    \
     2
    /
   3
```

输出: [1,3,2]
进阶:  递归算法很简单，你可以通过迭代算法完成吗？

## 伪代码

```
栈 stack
current = root

while(current 不为 null 或 stack 不为空) {
   while(current) {
      current 入栈
      current = current 的左子节点
   }
   current = stack 顶帧出栈
   访问 current
   current = current 的右子节点
}
```

## References

- [leetcode-cn](https://leetcode-cn.com/problems/binary-tree-inorder-traversal)
