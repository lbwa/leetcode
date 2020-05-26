export function frequencySort(string: string) {
  const wordsMap = createWordsMap(string)
  const queue = new PriorityQueue<string>((a, b) => {
    // 以下 if 语句用于
    // 1. 频率相等时，小字母在前
    // 2，始终将相同频率的相同字母紧靠在一起输出
    if (wordsMap[a] === wordsMap[b]) {
      return a < b
    }
    return wordsMap[a] > wordsMap[b]
  })

  for (const char of string) {
    queue.insert(char)
  }

  let i = string.length
  const result: string[] = []
  while (i--) {
    result.push(queue.poll())
  }
  return result.join('')
}

function createWordsMap(string: string) {
  const map: Record<string, number> = {}
  for (const char of string) {
    if (map[char] !== undefined) {
      map[char] += 1
    } else {
      map[char] = 0
    }
  }
  return map
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

  swim(index: number) {
    while (
      index > 1 &&
      this.comparator(this.heap[index] as E, this.heap[index >> 1] as E)
    ) {
      swap(this.heap, index, index >> 1)
      index = index >> 1
    }
  }

  sink(index: number) {
    while (2 * index <= this.size) {
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
    const deleted = this.heap[1]
    swap(this.heap, 1, this.size)
    this.heap.length--
    this.sink(1)
    return deleted
  }
}
