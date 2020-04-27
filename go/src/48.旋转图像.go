/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
// 方法1：先转置再逐行翻转
// 方法2：每次旋转4个元素
func rotate(matrix [][]int)  {
	n := len(matrix)
	for i := 0; i < n / 2 ; i ++ {
		for j := i; j < n - 1 - i; j++ {
			matrix[i][j], matrix[j][n - i - 1], matrix[n - i - 1][n - j - 1], matrix[n - j - 1][i] = 
				matrix[n - j - 1][i], matrix[i][j], matrix[j][n - i - 1], matrix[n - i - 1][n - j - 1]
		}
	}
}

func rotate2(matrix [][]int)  {
	n := len(matrix)
	for i := 0; i < (n + 1) / 2 ; i ++ {
		for j := 0; j < n / 2; j++ {
			matrix[i][j], matrix[j][n - i - 1], matrix[n - i - 1][n - j - 1], matrix[n - j - 1][i] = 
				matrix[n - j - 1][i], matrix[i][j], matrix[j][n - i - 1], matrix[n - i - 1][n - j - 1]
		}
	}
}
// @lc code=end

