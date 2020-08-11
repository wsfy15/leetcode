/*
 * @lc app=leetcode.cn id=1389 lang=golang
 *
 * [1389] 按既定顺序创建目标数组
 */

// @lc code=start
func createTargetArray(nums []int, index []int) []int {
	n := len(nums)
	res := make([]int, n)
	for k, i := range index {
		copy(res[i+1:], res[i:])
		res[i] = nums[k]
	}

	return res
}

// @lc code=end

