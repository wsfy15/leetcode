/*
 * @lc app=leetcode.cn id=188 lang=rust
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
impl Solution {
    pub fn max_profit(k: i32, prices: Vec<i32>) -> i32 {
        let n = prices.len();
        let limit = k as usize;

        // dp[i][j]: i表示经过了i次买入，j为0表示没有持有，j为1表示持有股票
        let mut dp = vec![vec![0; 2]; limit + 1];

        for i in 0..=limit {
            dp[i][1] = -prices[0];
        }

        for day in 1..n {
            for i in 1..=limit {
                dp[i][0] = dp[i][0].max(dp[i][1] + prices[day]);
                dp[i][1] = dp[i][1].max(dp[i - 1][0] - prices[day]);
            }
        }

        dp[limit][0]
    }
}
// @lc code=end
