/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	power := 1
	powers := []int{}
	for power * power <= n {
		powers = append(powers, power * power)
		power++
	}

	dp := make([]int, n + 1)
	dp[1] = 1
	powerIndex, count := 0, len(powers)
	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt32
		for powerIndex + 1 < count && powers[powerIndex + 1] <= i {
			powerIndex++
		}

		for j := powerIndex; j >= 0; j-- {
			dp[i] = min(dp[i], dp[i-powers[j]] + 1)
		}
	}
	return dp[n]
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
// @lc code=end

