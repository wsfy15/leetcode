/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
// 反转一半的数字
func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x > 0) {
		return false
	}

	num := 0
	for x > num {
		num = num * 10 + x % 10
		x /= 10
	}

	// 对于奇数长度的数字 num/10 可以去掉最中间的数字
	return num == x || num / 10 == x
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	tmp, num := x, 0
	for tmp > 0 {
		num = num * 10 + tmp % 10
		tmp /= 10
	}
	return num == x
}
// @lc code=end

