/*
 * @lc app=leetcode.cn id=330 lang=golang
 *
 * [330] 按要求补齐数组
 */

// @lc code=start
func minPatches(nums []int, n int) int {
	// 需要添加数字个数、待检查数字、nums下标
	res, cur, index := 0, 1, 0
	for cur <= n {
		if index >= len(nums) {
			cur += cur
			res++
		} else {
			if cur >= nums[index] {
				cur += nums[index]
				index++
			} else {
				cur += cur
				res++
			}
		}
	}

	return res
}

// @lc code=end

