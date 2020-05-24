import { greaterThan, swap } from 'src/shared'

/**
 * 使用优先队列的优势在于，不用排序整个数组就可得到极值项，即空间优势
 */
export class PriorityQueue<E extends number> {
  public heap: [null, ...E[]] = [null]

  constructor(private comparator: (low: E, high: E) => boolean) {}

  get size() {
    return this.heap.length - 1
  }

  private swim(index: number) {
    while (
      index > 1 &&
      // 二叉堆（完全二叉树）中第 k 项的父元素为 k >> 1
      this.comparator(this.heap[index] as E, this.heap[index >> 1] as E)
    ) {
      swap(this.heap, index, index >> 1)
      index = index >> 1
    }
  }

  private sink(index: number) {
    while (2 * index <= this.size) {
      // 二叉堆（完全二叉树）中第 k 项的子元素为 2 * k 和 2 * k + 1
      let child = 2 * index
      if (
        child < this.size &&
        this.comparator(this.heap[child + 1] as E, this.heap[child] as E)
      ) {
        child++
      }
      if (!this.comparator(this.heap[child] as E, this.heap[index] as E)) break
      swap(this.heap, index, child)
      index = child
    }
  }

  insert(el: E) {
    const newIndex = this.size + 1
    this.heap[newIndex] = el
    this.swim(newIndex)
  }

  poll() {
    if (this.size <= 0) return

    const deleted = this.heap[1]
    swap(this.heap, 1, this.size)
    this.heap.length -= 1
    this.sink(1)

    return deleted
  }
}

function createNumsMap<E extends number>(nums: E[]) {
  return nums.reduce((map, current) => {
    if (map[current] !== undefined) {
      map[current] += 1
    } else {
      map[current] = 1
    }
    return map
  }, {} as Record<string, number>)
}

export function topKFrequent<E extends number>(nums: E[], k: number) {
  const frequencyMap = createNumsMap(nums)

  const queue = new PriorityQueue((current, pivot) =>
    greaterThan(frequencyMap[current], frequencyMap[pivot])
  )
  Object.keys(frequencyMap).forEach((key) => queue.insert(+key as any))

  const answer: E[] = []
  for (let i = k; i > 0; i--) {
    const deleted = queue.poll() as E
    if (deleted !== undefined && deleted !== null) {
      answer.push(deleted)
    }
  }
  return answer
}
