package medium973

import (
	"container/heap"
	"sort"
)

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(a, b int) bool {
		// 求解平方根时需要引入 float 类型，故存在误差，但仅求解平方和，即为 int 类型，故不用考虑误差
		spaceA := points[a][0]*points[a][0] + points[a][1]*points[a][1]
		spaceB := points[b][0]*points[b][0] + points[b][1]*points[b][1]
		return spaceA < spaceB
	})

	return points[:K]
}

type point struct {
	distance     int
	coordination []int
}

type pq []point

func (q *pq) Len() int             { return len(*q) }
func (q *pq) Less(i, j int) bool   { return (*q)[i].distance > (*q)[j].distance }
func (q *pq) Swap(i, j int)        { (*q)[i], (*q)[j] = (*q)[j], (*q)[i] }
func (q *pq) Push(val interface{}) { *q = append(*q, val.(point)) }
func (q *pq) Pop() interface{} {
	queue := *q
	deleted := queue[len(queue)-1]
	*q = queue[:len(queue)-1]
	return deleted
}

func kClosestPQ(points [][]int, k int) (ans [][]int) {
	h := make(pq, k) // 控制大小为 k，即控制二叉堆大小为 k

	// 用第 0~k-1 的 k 个点填充堆容器
	for i, po := range points[:k] {
		h[i] = point{po[0]*po[0] + po[1]*po[1], po}
	}

	heap.Init(&h) // O(n) 初始化二叉堆

	for _, po := range points[k:] {
		// 与堆顶比较，即 h[0]
		if distance := po[0]*po[0] + po[1]*po[1]; distance < h[0].distance {
			h[0] = point{distance, po}
			// 在更新索引 i 的值后，通过 Fix 方法重建堆
			heap.Fix(&h, 0) // 此 if 语句执行效果即堆的插入，执行效率比 pop 后 push 要快
		}
	}

	// 当线性扫描完成时，大小为 k 的二叉堆中留下的点即为最接近原点的 k 个点
	for _, po := range h {
		ans = append(ans, po.coordination)
	}
	return
}
