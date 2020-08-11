/*
 * @lc app=leetcode.cn id=1342 lang=golang
 *
 * [1342] 将数字变成 0 的操作次数
 */

// @lc code=start
func numberOfSteps(num int) int {
	res := 0
	for num > 0 {
		res++
		if num&1 == 1 {
			num--
		} else {
			num >>= 1
		}
	}

	return res
}

// @lc code=end

