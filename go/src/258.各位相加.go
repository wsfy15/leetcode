/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
// 令y为num个位上的数，x为num/10
// f(x*10 + y) = f(x*9 + x + y) = f(x + y) 因此 f(t) = t % 9
func addDigits(num int) int {
	if num > 9 {
		num = num % 9
		if num == 0 {
			return 9
		}
	}

	return num
}
// @lc code=end

