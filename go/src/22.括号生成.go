/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	left, right := 1, 0
	cur := "("
	res := dfs(cur, left, right, n)

	return res
}

func dfs(cur string, left, right, n int) []string {
	if left == n && right == n {
		return []string{cur}
	}

	var res []string
	if left < n {
		res = append(res, dfs(cur + "(", left + 1, right, n)...)
	}
	if left > right {
		res = append(res, dfs(cur + ")", left, right + 1, n)...)
	}

	return res
}
// @lc code=end

