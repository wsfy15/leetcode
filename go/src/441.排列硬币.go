/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

// @lc code=start
func arrangeCoins(n int) int {
	return int(math.Sqrt(2.0) * math.Sqrt(float64(n) + 0.125) - 0.5)
}
// @lc code=end

