/*
 * @lc app=leetcode.cn id=213 lang=rust
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        if nums.len() == 1 {
            return nums[0];
        }

        return Self::util(nums[1..].to_vec()).max(Self::util(nums[..nums.len() - 1].to_vec()));
    }

    fn util(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        if n == 0 {
            return 0;
        }

        let mut dp = vec![0; n + 1];
        dp[1] = nums[0];

        nums.iter().enumerate().skip(1).for_each(|(i, x)| {
            dp[i + 1] = dp[i].max(dp[i - 1] + x);
        });

        dp[n]
    }
}
// @lc code=end
