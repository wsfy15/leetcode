/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */

// @lc code=start
func getSum(a int, b int) int {
	carry := a & b
	a = a ^ b
	
	for carry > 0 {
		a, carry = a ^ (carry << 1), a & (carry << 1)
	}

	return a
}
// @lc code=end

