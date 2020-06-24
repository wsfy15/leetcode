/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] 子数组最大平均数 I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	max, sum := math.MinInt32, 0
	n := len(nums)
	for i := 0; i < k - 1; i++ {
		sum += nums[i]
	}

	for j := k - 1; j < n; j++ {
		sum += nums[j]
		if sum > max {
			max = sum
		}
		sum -= nums[j - k + 1]
	}

	return float64(max) / float64(k)
}
// @lc code=end

