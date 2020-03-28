/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	rows := len(triangle)
	if rows == 0 {
		return 0
	}
	cols := len(triangle[rows-1])
	states := make([]int, cols)
	states[0] = triangle[0][0]
	for i := 1; i < rows; i++ {
		// 必须从右往左计算，因为从左往右计算会影响后面的计算结果
		// 最右的元素
		curCols := len(triangle[i])
		states[curCols - 1] = states[curCols - 2] + triangle[i][curCols - 1]
		for j := len(triangle[i]) - 2; j > 0; j-- {
			min := states[j]
			if states[j - 1] < min {
				min = states[j - 1]
			}
			states[j] = min + triangle[i][j]
		}
		states[0] += triangle[i][0]
	}

	min := states[0]
	for i := 1; i < cols; i++ {
		if states[i] < min {
			min = states[i]
		}
	}

	return min
}
// @lc code=end

