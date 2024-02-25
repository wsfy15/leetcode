/*
 * @lc app=leetcode.cn id=714 lang=rust
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
impl Solution {
    pub fn max_profit(prices: Vec<i32>, fee: i32) -> i32 {
        let (mut hold, mut not_hold) = (-prices[0] - fee, 0);
        prices.iter().skip(1).for_each(|price| {
            not_hold = not_hold.max(hold + price);
            hold = hold.max(not_hold - fee - price);
        });
        not_hold
    }
}
// @lc code=end
