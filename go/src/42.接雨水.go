/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	// 单调栈 https://oi-wiki.org/ds/monotonous-stack/
	stack := []int{}
	res := 0
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack) - 1]] {
			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			if len(stack) > 0 {
				newTop := stack[len(stack) - 1]
				distance := i - newTop - 1
				diffHeight := height[i]
				if height[i] > height[newTop] {
					diffHeight = height[newTop]
				}
				diffHeight -= height[top]
				res += distance * diffHeight
			}
		}

		stack = append(stack, i)
	}

	return res
}
// @lc code=end

