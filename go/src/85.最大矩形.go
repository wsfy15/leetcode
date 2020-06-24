/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	heights := make([]int, n + 2) // heights[0] heights[n + 1] 都是0，用于 getMaxArea
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j+1]++
			} else {
				heights[j+1] = 0
			}
		}
		res = max(res, getMaxArea(heights))
	}

	return res
}

func getMaxArea(heights []int) int {
	stack := []int{} // 递增栈
	area := 0
	for i, v := range heights {
		n := len(stack) - 1
		for n > 0 && heights[stack[n]] > v {
			area = max(area, (i - stack[n - 1] - 1) * heights[stack[n]])
			stack = stack[:n]
			n--
		}

		stack = append(stack, i)
	}

	return area
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
// @lc code=end

