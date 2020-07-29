/*
 * @lc app=leetcode.cn id=821 lang=golang
 *
 * [821] 字符的最短距离
 */

// @lc code=start
func shortestToChar(S string, C byte) []int {
	n := len(S)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		if S[i] == C {
			// 往前扫描
			for j := i - 1; j >= 0; j-- {
				if (S[j] != C && res[j] == 0) || res[j] > i - j {
					res[j] = i - j
				} else {
					break
				}
			}

			// 往后扫描
			j := i + 1
			for ; j < n && S[j] != C; j++ {
				res[j] = j - i
			}
			i = j - 1
		}
	}

	return res
}
// @lc code=end

