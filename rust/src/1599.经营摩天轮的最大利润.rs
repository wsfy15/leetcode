/*
 * @lc app=leetcode.cn id=1599 lang=rust
 *
 * [1599] 经营摩天轮的最大利润
 */

// @lc code=start
impl Solution {
    pub fn min_operations_max_profit(
        customers: Vec<i32>,
        boarding_cost: i32,
        running_cost: i32,
    ) -> i32 {
        let mut run_times = -1;
        let n = customers.len();
        let (mut profit, mut max_profit) = (0, 0);
        let (mut i, mut wait) = (0, 0);
        while i < n || wait > 0 {
            if i < n {
                wait += customers[i];
            }

            let count = wait.min(4);
            wait -= count;
            profit += boarding_cost * count - running_cost;
            i += 1;
            if profit > max_profit {
                max_profit = profit;
                run_times = i as i32;
            }
        }

        run_times
    }
}
// @lc code=end
