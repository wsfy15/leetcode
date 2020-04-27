/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	n, end := len(nums), 0
	for i := 0; i <= end; i++ {
		if i + nums[i] > end {
			end = i + nums[i]
		}
		if end >= n - 1 {
			return true
		}
	}

	return false
}
// @lc code=end

