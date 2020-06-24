/*
 * @lc app=leetcode.cn id=784 lang=golang
 *
 * [784] 字母大小写全排列
 */

// @lc code=start
func letterCasePermutation(S string) []string {
	var res []string
	dfs(S, &res, 0, "")
	return res
}

func dfs(S string, res *[]string, index int, cur string) {
	if index == len(S) {
		*res = append(*res, cur)
		return
	}

	if 'a' <= S[index] && S[index] <= 'z' {
		dfs(S, res, index + 1, cur + string(S[index]))
		dfs(S, res, index + 1, cur + string(S[index] - 32))
	} else if 'A' <= S[index] && S[index] <= 'Z' {
		dfs(S, res, index + 1, cur + string(S[index]))
		dfs(S, res, index + 1, cur + string(S[index] + 32))
	} else {
		dfs(S, res, index + 1, cur + string(S[index]))
	}
}
// @lc code=end

