/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(num int) bool {
	if num < 0 || num & (num - 1) > 0 {
		return false
	}
	return num & 0x55555555 > 0
	// return num % 3 == 1 
	// 如果是4的幂次，即num = (3 + 1) ^ n，(3 + 1) * (3 + 1) * …… * (3 + 1) 除了结尾的一个1，其他都是3的倍数
}
// @lc code=end

