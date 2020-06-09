import { BinaryTreeNode } from 'data-structures/binary-tree'

/**
 * 一棵有效的二叉搜索树的左子树的所有节点均大于当前节点，而右子树的所有节点均小于当前节点
 *
 * @core 递归法的本质是确定验证整个树中的节点是否满足一定的值的范围
 */
export function isValidBST(node: BinaryTreeNode<number> | null) {
  return validator(node, -Infinity, Infinity)
}

function validator(
  node: BinaryTreeNode<number> | null,
  lower: number,
  higher: number
): boolean {
  if (node === null || parent === null) return true

  /**
   * 1. 当验证左子树时，lower 不为 Infinity，那么验证当前节点值，是否小于父节点的值
   * 当不大于父节点的值时，即不满足二叉搜索树的定义，故返回 false
   *
   * 2. 右子树反之
   */
  if (node.value <= lower || node.value >= higher) return false

  return (
    // 验证左子树节点的值是否大于下界，小于上界，即左子树的节点要小于当前节点的值
    validator(node.left, lower, node.value) &&
    // 验证右子树的值是否大于下届，小于上界，即右子树的节点的值要始终大于当前节点的值
    validator(node.right, node.value, higher)
  )
}
