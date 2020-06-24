/*
 * @lc app=leetcode.cn id=693 lang=golang
 *
 * [693] 交替位二进制数
 */

// @lc code=start
func hasAlternatingBits(n int) bool {
	all1 := n ^ (n >> 1) // 如果是交替位，那么n右移一位后与n异或 就是全1
	return all1 & (all1 + 1) == 0 // all1 + 1就是2的幂次
}
// @lc code=end

