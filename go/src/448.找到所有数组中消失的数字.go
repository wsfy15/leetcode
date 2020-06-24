/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[abs(nums[i]) - 1] = -abs(nums[abs(nums[i]) - 1])
	}

	res := []int{}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			res = append(res, i + 1)
		}
	}

	return res
}

func abs(n int) int {
	y := n >> 31          // y ← x >> 31
	return (n ^ y) - y    // (x ⨁ y) - y
}
// @lc code=end

