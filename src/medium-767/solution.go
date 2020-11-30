package medium767

import (
	"container/heap"
	"sort"
)

var counts [26]int

// 结构体中的匿名字段，由类型名取该字段值，如下使用 `IntSlice` 取值
// https://stackoverflow.com/questions/28014591/nameless-fields-in-go-structs
type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool {
	return counts[h.IntSlice[i]] > counts[h.IntSlice[j]]
}
func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	deleted := a[len(a)-1]
	h.IntSlice = h.IntSlice[:len(a)-1]
	return deleted
}
func (h *hp) push(v int) {
	heap.Push(h, v)
}
func (h *hp) pop() int {
	return heap.Pop(h).(int)
}

func reorganizeString(s string) string {
	n := len(s)
	if n <= 1 {
		return ""
	}

	counts = [26]int{}
	// create counts
	var maxCounts int
	for _, ch := range s {
		ch -= 'a'
		counts[ch]++
		if maxCounts < counts[ch] {
			maxCounts = counts[ch]
		}
	}

	if maxCounts > (n+1)/2 {
		return ""
	}

	// init heap
	h := &hp{}
	for i, count := range counts[:] {
		if count > 0 {
			h.IntSlice = append(h.IntSlice, i)
		}
	}
	heap.Init(h)

	// construct char code slice
	ans := make([]byte, 0, n)
	for len(h.IntSlice) > 1 {
		i, j := h.pop(), h.pop()
		ans = append(ans, byte('a'+i), byte('a'+j))
		if counts[i]--; counts[i] > 0 {
			h.push(i)
		}
		if counts[j]--; counts[j] > 0 {
			h.push(j)
		}
	}
	if len(h.IntSlice) > 0 {
		ans = append(ans, byte('a'+h.IntSlice[0]))
	}
	return string(ans)
}
