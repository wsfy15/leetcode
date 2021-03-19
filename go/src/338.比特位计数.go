/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(num int) []int {
	states := make([]int, num + 1)
	for i := 1; i <= num; i++ {
		states[i] = states[i >> 1]
		if i & 1 == 1 {
			states[i]++
		}
	}
	return states
}
// @lc code=end

