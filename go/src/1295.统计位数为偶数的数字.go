/*
 * @lc app=leetcode.cn id=1295 lang=golang
 *
 * [1295] 统计位数为偶数的数字
 */

// @lc code=start
func findNumbers(nums []int) int {
	res := 0
	for _, num := range nums {
		if (num > 9 && num < 100) || (num > 999 && num < 10000) {
			res++
		}
	}
	return res
}

// @lc code=end

