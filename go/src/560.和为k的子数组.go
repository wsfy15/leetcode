/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	// 前缀和
	n := len(nums)
	sum, res := 0, 0
	counter := map[int]int{}
	counter[0] = 1
	for i := 0; i < n; i++ {
		sum += nums[i]
		if v, ok := counter[sum - k]; ok {
			res += v
		}
		counter[sum]++
	}
	return res
}
// @lc code=end

