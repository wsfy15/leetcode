/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 */

// @lc code=start
func titleToNumber(s string) int {
	res, n := 0, len(s)
	for i := 0; i < n; i++ {
		res = res * 26 + int(s[i] - 'A' + 1)
	}

	return res
}
// @lc code=end

