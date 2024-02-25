/*
 * @lc app=leetcode.cn id=122 lang=rust
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let mut res = 0;
        let mut last_price = prices[0];
        prices.iter().skip(1).for_each(|&price| {
            res += 0.max(price - last_price);
            last_price = price;
        });

        res
    }
}
// @lc code=end
