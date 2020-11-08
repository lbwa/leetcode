package easy559

type node struct {
	Val      int
	Children []*node
}

func maxDepthBfs(root *node) (depth int) {
	if root == nil {
		return
	}
	queue := []*node{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		depth++
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			if len(current.Children) > 0 {
				queue = append(queue, current.Children...)
			}
		}
	}
	return
}

func maxDepthDfs(root *node) int {
	var max func(int, int) int
	var dfs func(*node) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dfs = func(root *node) int {
		if root == nil {
			return 0
		}
		var depth int
		for _, child := range root.Children {
			depth = max(depth, dfs(child))
		}
		return depth + 1
	}
	return dfs(root)
}
