/*
 * @lc app=leetcode.cn id=1403 lang=golang
 *
 * [1403] 非递增顺序的最小子序列
 */

// @lc code=start
func minSubsequence(nums []int) []int {
	total := 0
	for _, num := range nums {
		total += num
	}

	res := []int{}
	cur := 0
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		cur += nums[i]
		res = append(res, nums[i])
		if cur > total-cur {
			break
		}
	}

	return res
}

// @lc code=end

