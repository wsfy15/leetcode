/*
 * @lc app=leetcode.cn id=1029 lang=golang
 *
 * [1029] 两地调度
 */

// @lc code=start
func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0] - costs[i][1] < costs[j][0] - costs[j][1]
	})

	n := len(costs) / 2
	res := 0
	for i := 0; i < n; i++ {
		res += costs[i][0]
	}
	for i := n; i < len(costs); i++ {
		res += costs[i][1]
	}
	return res
}
// @lc code=end

