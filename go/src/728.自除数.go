/*
 * @lc app=leetcode.cn id=728 lang=golang
 *
 * [728] 自除数
 */

// @lc code=start
func selfDividingNumbers(left int, right int) []int {
	res := []int{}
	for i := left; i <= right; i++ {
		if isValid(i) {
			res = append(res, i)
		}
	}
	return res
}

func isValid(num int) bool {
	tmp := num
	for num > 0 {
		digit := num % 10
		if digit == 0 || tmp % digit != 0 {
			return false
		}
		num /= 10
	}
	return true
}
// @lc code=end

