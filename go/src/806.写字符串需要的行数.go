/*
 * @lc app=leetcode.cn id=806 lang=golang
 *
 * [806] 写字符串需要的行数
 */

// @lc code=start
func numberOfLines(widths []int, S string) []int {
	cur, res := 100, 0

	for i := 0; i < len(S); i++ {
		if widths[S[i] - 'a'] + cur > 100 {
			res++
			cur = 0
		} 
		cur += widths[S[i] - 'a']
	}
	return []int{res, cur}
}
// @lc code=end

