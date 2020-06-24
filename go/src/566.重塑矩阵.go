/*
 * @lc app=leetcode.cn id=566 lang=golang
 *
 * [566] 重塑矩阵
 */

// @lc code=start
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	m, n := len(nums), len(nums[0])
	if m * n != r * c {
		return nums
	}

	row, col := 0, 0

	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
		for k := 0; k < c; k++ {
			res[i][k] = nums[row][col]
			col++
			if col == n {
				col = 0
				row++
			}
		}
	}

	return res
}
// @lc code=end

