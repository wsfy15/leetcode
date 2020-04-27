/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	var num int64 = int64(x)
	var cur int64
	var flag int64 = 1
	if num < 0 {
		num = -num
		flag = -1
	}

	for num > 0 {
		cur = cur * 10 + num % 10
		num /= 10
	}
	cur *= flag

	if cur > math.MaxInt32 || cur < math.MinInt32 {
		return 0
	}
	return int(cur)
}
// @lc code=end

