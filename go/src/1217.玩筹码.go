/*
 * @lc app=leetcode.cn id=1217 lang=golang
 *
 * [1217] 玩筹码
 */

// @lc code=start
func minCostToMoveChips(chips []int) int {
	even, odd := 0, 0
	for _, num := range chips {
		if num&1 == 1 {
			odd++
		} else {
			even++
		}
	}
	if odd < even {
		return odd
	}
	return even
}

// @lc code=end

