/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	num, count := x ^ y, 0
	for num > 0 {
		num &= num - 1
		count++
	}
	return count
}
// @lc code=end

