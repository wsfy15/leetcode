/*
 * @lc app=leetcode.cn id=2706 lang=rust
 *
 * [2706] 购买两块巧克力
 */

// @lc code=start
impl Solution {
    pub fn buy_choco(prices: Vec<i32>, money: i32) -> i32 {
        let mut prices = prices;
        prices.sort_unstable();
        let cost = prices[0] + prices[1];
        if cost > money {
            return money;
        }
        money - cost
    }
}
// @lc code=end
