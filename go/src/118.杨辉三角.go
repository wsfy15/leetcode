/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	res := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		cur := make([]int, i + 1)
		cur[0], cur[i] = 1, 1
		for j := 1; j < i; j++ {
			cur[j] = res[i - 1][j - 1] + res[i - 1][j]
		}

		res = append(res, cur)
	}

	return res
}
// @lc code=end

