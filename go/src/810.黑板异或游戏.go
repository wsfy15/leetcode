/*
 * @lc app=leetcode.cn id=810 lang=golang
 *
 * [810] 黑板异或游戏
 */

// @lc code=start
func xorGame(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum ^= num
	}

	return sum == 0 || len(nums)%2 == 0
}

// @lc code=end

