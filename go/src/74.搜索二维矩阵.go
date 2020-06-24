/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 || target < matrix[0][0] || target > matrix[m - 1][n - 1] {
		return false
	}

	low, high := 0, m - 1
	for low < high {
		mid := low + (high - low) >> 1

		if matrix[mid][0] <= target && (mid == m - 1 || matrix[mid + 1][0] > target) {
			low = mid
			break;
		} else {
			if matrix[mid][0] < target {
				low = mid + 1
			} else {
				high = mid
			}
		}
	}

	row := low
	low, high = 0, n - 1
	for low < high {
		mid := low + (high - low) >> 1

		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return matrix[row][low] == target
}
// @lc code=end

