/*
 * @lc app=leetcode.cn id=198 lang=rust
 *
 * [198] 打家劫舍
 */

// @lc code=start
impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut dp = vec![0; n + 1];
        dp[1] = nums[0];

        nums.iter().enumerate().skip(1).for_each(|(i, x)| {
            dp[i + 1] = dp[i].max(dp[i - 1] + x);
        });

        dp[n]
    }
}
// @lc code=end
