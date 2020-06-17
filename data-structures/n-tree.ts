export class NTreeNode<V> {
  constructor(public value: V, public children: NTreeNode<V>[] = []) {}
}
