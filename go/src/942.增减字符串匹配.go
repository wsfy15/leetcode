/*
 * @lc app=leetcode.cn id=942 lang=golang
 *
 * [942] 增减字符串匹配
 */

// @lc code=start
func diStringMatch(S string) []int {
	n := len(S)
	low, high := 0, n // 每次选择的，要么是最小的，要么是最大的
	res := make([]int, n + 1)

	for i := 0; i < n; i++ {
		if S[i] == 'I' {
			res[i] = low
			low++
		} else {
			res[i] = high
			high--
		}
	}
	res[n] = low
	return res
}
// @lc code=end

