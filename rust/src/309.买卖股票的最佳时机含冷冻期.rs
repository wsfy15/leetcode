/*
 * @lc app=leetcode.cn id=309 lang=rust
 *
 * [309] 买卖股票的最佳时机含冷冻期
 */

// @lc code=start
impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let n = prices.len();
        if n < 2 {
            return 0;
        }

        // dp[i][0]：第i天没有持股且当天没有卖出的最大收益
        // dp[i][1]：第i天持股的最大收益
        // dp[i][2]：第i天没有持股且当天卖出的最大收益
        let mut dp = vec![vec![0; 3]; n];
        dp[0][1] = -prices[0];

        for i in 1..n {
            dp[i][0] = dp[i - 1][0].max(dp[i - 1][2]);
            dp[i][1] = dp[i - 1][1].max(dp[i - 1][0] - prices[i]);
            dp[i][2] = dp[i - 1][1] + prices[i];
        }

        dp[n - 1][0].max(dp[n - 1][2])
    }
}
// @lc code=end
