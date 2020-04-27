/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	left, right := 0, len(height) - 1
	res := 0
	for left < right {
		tmp := (right - left) * min(height[left], height[right])
		if tmp > res {
			res = tmp
		}
		
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

