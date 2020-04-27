/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	for i := 32; i > 0 ; i-- {
		res <<= 1
		res += num & 1
		num >>= 1
	}	
	return res
}
// @lc code=end

