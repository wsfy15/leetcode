/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	stack := []int{0} // 单调递增栈 存储index
	max := 0

	for i := 1; i < len(heights); i++ {
		j := len(stack) - 1
		for j > 0 && heights[i] < heights[stack[j]] {
			// 关键：要减去前一个位置的坐标，应对比heights[stack[j]]高的被pop的元素
			// [0, 2, 1, 2]： [0, 2] -> [0, (2被pop了) 1] -> [0, 1, 2] 
			area := (i -(stack[j - 1] + 1)) * heights[stack[j]]
			if area > max {
				max = area
			}
			stack = stack[:j]
			j--
		}
		stack = append(stack, i)
	}

	return max
}
// @lc code=end

