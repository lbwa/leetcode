import { swap } from 'src/shared'

export class KthLargest {
  private heap: [null, ...number[]] = [null]

  private maxSize: number

  get size() {
    return this.heap.length - 1
  }

  get peek() {
    return this.heap[1]
  }

  constructor(k: number, nums: number[]) {
    this.maxSize = k
    nums.forEach((num) => this.add(num))
  }

  private swim(index: number) {
    const { heap } = this
    while (
      index > 1 &&
      (heap[index] as number) < (heap[index >> 1] as number)
    ) {
      swap(heap, index, index >> 1)
      index = index >> 1
    }
  }

  private sink(index: number) {
    const { heap } = this
    while (2 * index <= this.size) {
      let child = 2 * index

      if (
        child < this.size &&
        (heap[child + 1] as number) < (heap[child] as number)
      ) {
        child++
      }

      // 当子节点不再大于当前节点时，那么取消下沉，否则继续下沉大节点，使得上层节点较小
      if (!((heap[index] as number) > (heap[child] as number))) break

      swap(heap, index, child)
      index = child
    }
  }

  private insert(el: number) {
    const newIndex = this.size + 1
    this.heap[newIndex] = el
    this.swim(newIndex) // 恢复“堆有序”
  }

  add(val: number): number {
    // 小根堆未满时，直接插入
    if (this.size < this.maxSize) {
      this.insert(val)
    } else if (
      /* 小根堆满时，且要插入的值大于根节点时，那么 poll 根节点，并插入新的节点值
       */ val > this.peek
    ) {
      this.poll()
      this.insert(val)
    }

    return this.peek
  }

  poll() {
    const deleted = this.peek
    swap(this.heap, 1, this.size)
    this.heap.length -= 1
    this.sink(1) // 必须在删除待删除节点之后，恢复“堆有序”
    return deleted
  }
}
