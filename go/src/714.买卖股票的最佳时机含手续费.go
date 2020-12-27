/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	hold, notHold := -fee-prices[0], 0
	for i := 1; i < len(prices); i++ {
		notHold = max(notHold, hold+prices[i])
		hold = max(notHold-fee-prices[i], hold)
	}

	return notHold
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

