/*
 * @lc app=leetcode.cn id=1304 lang=golang
 *
 * [1304] 和为零的N个唯一整数
 */

// @lc code=start
func sumZero(n int) []int {
	res := []int{}
	if n&1 == 1 {
		res = append(res, 0)
	}

	n >>= 1
	for n > 0 {
		res = append(res, n, -n)
		n--
	}

	return res
}

// @lc code=end

