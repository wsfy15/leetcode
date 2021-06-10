/*
 * @lc app=leetcode.cn id=523 lang=golang
 *
 * [523] 连续的子数组和
 */

// @lc code=start
func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	prefix := make([]int, n+1)
	m := map[int][]int{}
	m[0] = []int{0}
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
		m[prefix[i]%k] = append(m[prefix[i]%k], i)
	}

	for i := 1; i <= n; i++ {
		index := prefix[i] % k
		for _, v := range m[index] {
			if v < i-1 || v > i+1 {
				return true
			}
		}
	}

	return false
}

// @lc code=end

