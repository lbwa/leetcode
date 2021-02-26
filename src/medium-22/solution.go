package medium22

import "strings"

func generateParenthesis(n int) []string {
	ans := []string{}
	str := []string{}
	var dfs func(int, int)
	dfs = func(left, right int) {
		if len(str) == 2*n {
			ans = append(ans, strings.Join(str, ""))
			return
		}

		if left < n {
			str = append(str, "(")
			dfs(left+1, right)
			str = str[:len(str)-1]
		}
		if right < left {
			str = append(str, ")")
			dfs(left, right+1)
			str = str[:len(str)-1]
		}
	}

	dfs(0, 0)
	return ans
}

func generateParenthesis0(n int) (ans []string) {
	str := []string{}
	var dfs func(int, int)
	dfs = func(left, right int) {
		if left > n || (right > left) {
			return
		}
		if len(str) == 2*n {
			ans = append(ans, strings.Join(str, ""))
			return
		}

		// for 循环选择列表为 []{"(", ")"}
		str = append(str, "(")
		dfs(left+1, right)
		str = str[:len(str)-1]

		str = append(str, ")")
		dfs(left, right+1)
		str = str[:len(str)-1]
	}
	dfs(0, 0)
	return
}
