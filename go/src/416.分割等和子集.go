/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition2(nums []int) bool {
	n, sum := len(nums), 0
	for _, num := range nums {
		sum += num
	}

	if sum&1 == 1 {
		return false
	}

	sum >>= 1
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, sum+1)
	}

	dp[0][0] = true // 虽然选中nums[0]后不等于0，但方便了后面 j-nums[i] == 0 的情况

	if nums[0] <= sum {
		dp[0][nums[0]] = true
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= sum; j++ {
			dp[i][j] = dp[i-1][j]
			// if nums[i] == j {
			// 	dp[i][j] = true
			// } else if nums[i] < j {
			// 	dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			// }
			if nums[i] <= j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
		if dp[i][sum] {
			return true
		}
	}

	return dp[n-1][sum]
}

func canPartition(nums []int) bool {
	n, sum := len(nums), 0
	for _, num := range nums {
		sum += num
	}

	if sum&1 == 1 {
		return false
	}

	sum >>= 1
	dp := make([]bool, sum+1)
	if nums[0] <= sum {
		dp[nums[0]] = true
	}

	for i := 1; i < n; i++ {
		for j := sum; j >= nums[i]; j-- {
			if dp[sum] {
				return true
			}
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[sum]
}

// @lc code=end

