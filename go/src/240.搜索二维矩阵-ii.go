/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	// 分治法：
	// 左下角的元素是最左列的最大值，最下行的最小值。
	// 如果该元素大于target，目标矩阵去掉最下面一行；如果该元素小于target，目标矩阵去掉最左一行
	// 如果该元素等于target，返回true
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	return divide(matrix, target, len(matrix) - 1, 0)
}

func divide(matrix [][]int, target, biggestRow, smallestCol int) bool {
	if biggestRow < 0 || smallestCol >= len(matrix[0]) {
		return false
	}
	if matrix[biggestRow][smallestCol] == target {
		return true
	} else if matrix[biggestRow][smallestCol] > target {
		return divide(matrix, target, biggestRow - 1, smallestCol)
	} else {
		return divide(matrix, target, biggestRow, smallestCol + 1)
	}
}
// @lc code=end

