/*
 * @lc app=leetcode.cn id=368 lang=golang
 *
 * [368] 最大整除子集
 */

// @lc code=start
func largestDivisibleSubset(nums []int) []int {
	n, maxIndex := len(nums), 0
	sort.Ints(nums)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > dp[maxIndex] {
			maxIndex = i
		}
	}

	count := dp[maxIndex]
	res := []int{}
	for i := maxIndex; i >= 0; i-- {
		if dp[i] == count && nums[maxIndex]%nums[i] == 0 {
			res = append(res, nums[i])
			count--
			maxIndex = i
		}
	}

	return res
}

func max(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end

