/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	
	res := 0
	minStock := prices[0] // 记录从开始到当前的最低股价
	for i := 1; i < len(prices); i++ {
		if minStock > prices[i] {
			minStock = prices[i]
		} else {
			// 计算卖出得到的收益
			if benefit := prices[i] - minStock; benefit > res {
				res = benefit
			}
		}
	}
	return res
}

// states[i] = max(states[i - 1] + prices[i] - prices[i-1], 0)
func maxProfit2(prices []int) int {
	size := len(prices)
	if size == 0 {
			return 0
	}
	states := make([]int, size)
	states[0] = 0
	for i := 1; i < size; i++ {
			if states[i-1] + prices[i] - prices[i-1] > 0 {
					states[i] = states[i-1] + prices[i] - prices[i-1]
			}
	}
	max := 0
	for i := range states {
			if states[i] > max {
					max = states[i]
			}
	}
	return max
}
// @lc code=end

