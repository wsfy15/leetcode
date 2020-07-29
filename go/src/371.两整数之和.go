/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */

// @lc code=start
func getSum(a int, b int) int {
	carry := (a & b) << 1
	a = a ^ b
	
	for carry != 0 { // 有可能是负数
		a, carry = a ^ carry, (a & carry) << 1
	}

	return a
}
// @lc code=end

