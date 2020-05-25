export function topKFrequent(words: string[], k: number) {
  const wordsMap = createWordsMap(words)

  const queue = new PriorityQueue((a: string, b: string) => {
    if (wordsMap[a] > wordsMap[b]) {
      return true
    }

    if (wordsMap[a] === wordsMap[b]) {
      // JS 中字母字符串比较按照 ASCII 码比较
      return a < b
    }
    return false
  })

  Object.keys(wordsMap).forEach((word) => queue.insert(word))

  const answer = []
  while (k--) {
    answer.push(queue.poll())
  }

  return answer
}

function createWordsMap(words: string[]) {
  return words.reduce((map, word) => {
    if (map[word] !== undefined) {
      map[word]++
    } else {
      map[word] = 1
    }
    return map
  }, {} as Record<string, number>)
}

function swap<E>(list: E[], a: number, b: number) {
  ;[list[a], list[b]] = [list[b], list[a]]
}

class PriorityQueue<E> {
  private heap: [null, ...E[]] = [null]
  constructor(private comparator: (a: E, b: E) => boolean) {}

  get size() {
    return this.heap.length - 1
  }

  sink(index: number) {
    const { heap } = this
    while (2 * index <= this.size) {
      // k 项子元素为 2 * k 和 2 * k + 1
      let child = 2 * index
      if (
        child < this.size &&
        this.comparator(heap[child + 1] as E, heap[child] as E)
      ) {
        child++
      }
      if (!this.comparator(heap[child] as E, heap[index] as E)) break
      swap(heap, index, child)
      index = child
    }
  }

  swim(index: number) {
    const { heap } = this
    while (
      index > 1 &&
      this.comparator(heap[index] as E, heap[index >> 1] as E)
    ) {
      const parent = index >> 1
      swap(heap, index, parent)
      index = parent
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
