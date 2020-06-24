/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	
	for num % 5 == 0 {
		num /= 5
	}
	for num % 3 == 0 {
		num /= 3
	}
	for num & 1 == 0 {
		num >>= 1
	}
	return num == 1
}
// @lc code=end

