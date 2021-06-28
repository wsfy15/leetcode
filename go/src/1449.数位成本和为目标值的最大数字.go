/*
 * @lc app=leetcode.cn id=1449 lang=golang
 *
 * [1449] 数位成本和为目标值的最大数字
 */

// @lc code=start
func largestNumber(cost []int, target int) string {
	dp := make([]int, target+1) // dp[i]: i成本最多可以换到的数 数量
	for i := 1; i <= target; i++ {
		dp[i] = math.MinInt32
		for _, v := range cost {
			if i >= v {
				dp[i] = max(dp[i], dp[i-v]+1)
			}
		}
	}

	if dp[target] <= 0 {
		return "0"
	}

	sb := strings.Builder{}
	for target > 0 {
		for i := 9; i >= 1; i-- {
			if target >= cost[i-1] && dp[target]-1 == dp[target-cost[i-1]] {
				sb.WriteByte('0' + byte(i))
				target -= cost[i-1]
				break
			}
		}
	}

	return sb.String()
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

