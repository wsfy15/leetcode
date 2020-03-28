/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}
	states := make([]int, size) // 以当前元素结束的最长上升子序列长度

	for i := range nums {
		states[i] = 0
		for j := i - 1; j >= 0; j-- {
			// 对前面比nums[i]小的数字 对应的最长上升子序列 加一
			if nums[i] > nums[j] && states[j] > states[i] {
				states[i] = states[j]
			}
		}
		states[i]++
	}

	max := 0
	for _, v := range states {
		if max < v {
			max = v
		}
	}
	return max
}
// @lc code=end

