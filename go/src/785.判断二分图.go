/*
 * @lc app=leetcode.cn id=785 lang=golang
 *
 * [785] 判断二分图
 */

// @lc code=start
func isBipartite(graph [][]int) bool {
	n := len(graph)
	group := make([]int, n)
	for i := 0; i < n; i++ {
		if group[i] == 0 && !dfs(graph, group, i, 1) {
			return false
		}
	}
	return true
}

func dfs(graph [][]int, group []int, index, val int) bool {
	group[index] = val
	for i := 0; i < len(graph[index]); i++ {
		if group[graph[index][i]] == val {
			return false
		}
		if group[graph[index][i]] == 0 && !dfs(graph, group, graph[index][i], -val) {
			return false
		}
	}
	return true
}
// @lc code=end

