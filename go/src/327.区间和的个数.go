/*
 * @lc app=leetcode.cn id=327 lang=golang
 *
 * [327] 区间和的个数
 */

// @lc code=start
func countRangeSum(nums []int, lower int, upper int) int {
	res := 0
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
		for j := 0; j < i; j++ {
			val := preSum[i] - preSum[j]
			if val >= lower && val <= upper {
				res++
			}
		}
	}
	return res
}

// @lc code=end

