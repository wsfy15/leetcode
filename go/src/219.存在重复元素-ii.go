/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	pos := map[int]int{}
	for i, v := range nums {
		if p, ok := pos[v]; ok {
			if i - p <= k {
				return true
			}
		}
		pos[v] = i
	}

	return false
}
// @lc code=end

