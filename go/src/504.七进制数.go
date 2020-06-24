/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] 七进制数
 */

// @lc code=start
func convertToBase7(num int) string {
	res, factor := 0, 1
	if num < 0 {
		num = -num
		factor = -1
	}

	for num > 0 {
		cur := num % 7
		num /= 7
		res += factor * cur 
		factor *= 10
	}

	return strconv.Itoa(res)
}
// @lc code=end

