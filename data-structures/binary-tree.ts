export class BinaryTreeNode<V> {
  constructor(
    public value: V,
    public left: BinaryTreeNode<V> | null = null,
    public right: BinaryTreeNode<V> | null = null
  ) {}
}
