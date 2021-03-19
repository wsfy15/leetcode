/*
 * @lc app=leetcode.cn id=766 lang=golang
 *
 * [766] 托普利茨矩阵
 */

// @lc code=start
func isToeplitzMatrix2(matrix [][]int) bool {
	m := map[int]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if v, ok := m[j-i]; ok {
				if matrix[i][j] != v {
					return false
				}
			} else {
				m[j-i] = matrix[i][j]
			}
		}
	}
	return true
}

// 判断前一行的除最后一个元素外的所有元素，是否等于下一行除第一个元素的所有元素
// 即类似平移下去
func isToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m-1; i++ {
		if !equal(matrix[i][:n-1], matrix[i+1][1:]) {
			return false
		}
	}

	return true
}

func equal(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// @lc code=end

