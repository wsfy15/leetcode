/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H 指数
 */

// @lc code=start
func hIndex(citations []int) int {
	n := len(citations)
	sort.Ints(citations)
	for i := 0; i < n; i++ {
		if n - i <= citations[i] {
			return n - i
		}
	}

	return 0
}
// @lc code=end

