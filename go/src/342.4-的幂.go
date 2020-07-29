/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(num int) bool {
	if num <= 0 || num & (num - 1) > 0 { // 已经验证只有1个1了
		return false
	}
	// 100、10000…… 4的幂只会出现一位1，且只在奇数位上出现
	return num & 0x55555555 == num
	// return num % 3 == 1 
	// 如果是4的幂次，即num = (3 + 1) ^ n，(3 + 1) * (3 + 1) * …… * (3 + 1) 除了结尾的一个1，其他都是3的倍数
}
// @lc code=end

