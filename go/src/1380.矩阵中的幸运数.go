/*
 * @lc app=leetcode.cn id=1380 lang=golang
 *
 * [1380] 矩阵中的幸运数
 */

// @lc code=start
func luckyNumbers(matrix [][]int) []int {
	res := []int{}
	for _, row := range matrix {
		minIndex := 0
		for i, v := range row {
			if v < row[minIndex] {
				minIndex = i
			}
		}

		i := 0
		for ; i < len(matrix); i++ {
			if matrix[i][minIndex] > row[minIndex] {
				break
			}
		}
		if i == len(matrix) {
			res = append(res, row[minIndex])
		}
	}

	return res
}

// @lc code=end

