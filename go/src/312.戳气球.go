/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */

// @lc code=start
func maxCoins(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([][]int, n + 2)
	for i := 0; i < n + 2; i++ {
		dp[i] = make([]int, n + 2)
	}
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	for l := 1; l <= n; l++ { // len
		for i := 1; i <= n - l + 1; i++ {
			j := i + l - 1
			for k := i; k <= j; k++ { // 以k为[i,j]中最后一个爆炸的
				dp[i][j] = max(dp[i][j], dp[i][k-1] + dp[k+1][j] + nums[i-1] * nums[k] * nums[j+1])
			}
		}
	}
	return dp[1][n]
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
// @lc code=end

