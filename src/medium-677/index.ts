export class MapSum {
  val = 0
  next = {} as Record<string, MapSum>

  insert(key: string, val: number) {
    let node = this as MapSum
    for (const char of key) {
      if (!node.next[char]) {
        node.next[char] = new MapSum()
      }
      node = node.next[char]
    }
    node.val = val
  }

  sum(prefix: string): number {
    let node = this as MapSum
    for (const char of prefix) {
      if (!node.next[char]) {
        return 0
      }
      node = node.next[char]
    }
    // 此时 node 为 以 prefix 为前缀的 trie tree，那么求 node 的所有子节点的 val 之和（通过 dfs）
    return dfs(node)
  }
}

function dfs(sum: MapSum): number {
  let answer = sum.val
  for (const char in sum.next) {
    answer += dfs(sum.next[char])
  }
  return answer
}
