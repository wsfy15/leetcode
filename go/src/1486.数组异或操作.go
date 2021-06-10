/*
 * @lc app=leetcode.cn id=1486 lang=golang
 *
 * [1486] 数组异或操作
 */

// @lc code=start
func xorOperation(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		res ^= start
		start += 2
	}

	return res
}

// @lc code=end

