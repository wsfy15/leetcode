/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */

// @lc code=start
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	res := []int{0, 1}
	for i := 2; i <= n; i++ {
		// 对res从后往前的每一个数，在第i位添加1构成新的数
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, res[j] + 1 << (i - 1))
		}
	}

	return res
}
// @lc code=end

