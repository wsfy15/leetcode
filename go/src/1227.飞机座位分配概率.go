/*
 * @lc app=leetcode.cn id=1227 lang=golang
 *
 * [1227] 飞机座位分配概率
 */

// @lc code=start
func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	}
	return 0.5
}

// @lc code=end

