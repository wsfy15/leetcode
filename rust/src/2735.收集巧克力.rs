/*
 * @lc app=leetcode.cn id=2735 lang=rust
 *
 * [2735] 收集巧克力
 */

// @lc code=start
impl Solution {
    pub fn min_cost(nums: Vec<i32>, x: i32) -> i64 {
        let n = nums.len();
        // dp[k] 统计操作 k 次的总成本
        let mut dp: Vec<i64> = (0..n).map(|i| i as i64 * x as i64).collect();
        for i in 0..n {
            // 子数组左端点
            let mut num = nums[i];
            for j in i..(n + i) {
                // 子数组右端点（把数组视作环形的）
                num = num.min(nums[j % n]);
                dp[j - i] += num as i64;
            }
        }

        *dp.iter().min().unwrap()
    }
}
// @lc code=end
