/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := rowIndex; j > 0; j-- {
			res[j] += res[j-1]
		}
	}

	return res
}

// @lc code=end
