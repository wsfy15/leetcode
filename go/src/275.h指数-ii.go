/*
 * @lc app=leetcode.cn id=275 lang=golang
 *
 * [275] H指数 II
 */

// @lc code=start
func hIndex(citations []int) int {
	n := len(citations)
	if n == 0 || citations[n - 1] == 0 {
		return 0
	}
	
	left, right := 0, n - 1
	for left < right {
		mid := left + (right - left) >> 1
		count := n - mid

		// 文章数 大于 引用次数，往文章数减少但引用次数增加方向变化
 		if citations[mid] < count {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return n - left
}
// @lc code=end

