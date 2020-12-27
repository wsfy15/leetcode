/*
 * @lc app=leetcode.cn id=861 lang=golang
 *
 * [861] 翻转矩阵后的得分
 */

// @lc code=start
func matrixScore(A [][]int) int {
	if len(A) == 0 {
		return 0
	}

	row, col := len(A), len(A[0])
	for i := 0; i < row; i++ {
		if A[i][0] == 0 {
			for j := 0; j < col; j++ {
				A[i][j] ^= 1
			}
		}
	}

	res := row * (1 << (col - 1))
	for j := 1; j < col; j++ {
		count := 0
		for i := 0; i < row; i++ {
			if A[i][j] == 1 {
				count++
			}
		}

		if count > row - count {
			res += count * (1 << (col - j - 1))
		} else {
			res += (row - count) * (1 << (col - j - 1))
		}
	}


	return res
}
// @lc code=end

