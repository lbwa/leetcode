import { LinkedListNode } from 'data-structures/linked-list'

/**
 * @see https://leetcode-cn.com/problems/reverse-linked-list/solution/fan-zhuan-lian-biao-by-leetcode/
 * 递归版本，其关键在于反向工作。假设列表的其余部分已经被反转，现在我该如何反转它前面的
 * 部分？假设列表为
 * N1 -> ... -> Nk-1 -> Nk -> Nk+1 <- ... Nm <- ∅
 * 若从节点 Nk+1 到 Nm 已经被反转，而我们正处于 Nk
 * 所以 Nk.next.next = Nk
 * 要小心的是 N1 的下一项为 ∅，如果忽略了这一点，反转后的链表中可能会产生循环。如果使用
 * 大小为 2 的链表测试代码，则可能会捕获此错误。
 */
export function reverseLinkedList<T>(
  head: LinkedListNode<T> | null
): LinkedListNode<T> | null {
  if (head === null || head.next === null) return head

  /**
   * reverseLinkedList(2 -> 3 -> 4 -> null)，即以下语句要实现将从下一项开始
   * 反转所有链表项，后文表示如何在栈帧中处理节点反转
   */
  const last = reverseLinkedList(head.next)

  /**
   * 在递归栈中实现反转链表项，可理解为我子节点下的所有节点都已经反转好了，现在就剩我和
   * 我的子节点 没有完成最后的反转了，所以反转一下我和我的子节点
   *
   * 1. head.next 表示需要被反转的起始项，那么 head.next.next 表示将起始项的下一项
   * 指向起始项的前一项，即 head 项
   * 2. 可结合 2 -> 3 -> 4 -> null 示例链表来思考以下语句
   */
  head.next.next = head // 即反转链接，即 3 不再指向 4，而是指向前一节点

  // 断开原有的项指向下一项的链接，因为在反转中，通过上一语句指向反转后的下一节点，即恢
  // 复链接
  // 该语句主要针对于原有的 head 节点，它在反转后为 tail 节点，那么其 next 应为 null
  head.next = null // 即 2 -> 3 -> 4 -> null 中 2 不再指向 3，而是 null
  return last
}
