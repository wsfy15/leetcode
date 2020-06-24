/*
 * @lc app=leetcode.cn id=766 lang=golang
 *
 * [766] 托普利茨矩阵
 */

// @lc code=start
func isToeplitzMatrix(matrix [][]int) bool {
	m := map[int]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if v, ok := m[j - i]; ok {
				if matrix[i][j] != v {
					return false
				}
			} else {
				m[j - i] = matrix[i][j]
			}
		}
	}
	return true
}
// @lc code=end

