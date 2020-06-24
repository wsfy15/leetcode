/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i, j := 0, 0
	n, m := len(g), len(s)
	count := 0
	for i < n && j < m {
		if s[j] >= g[i] {
			i++
			count++
		}
		j++
	}

	return count
}
// @lc code=end

