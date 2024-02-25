/*
 * @lc app=leetcode.cn id=689 lang=rust
 *
 * [689] 三个无重叠子数组的最大和
 */

// @lc code=start
impl Solution {
    pub fn max_sum_of_three_subarrays(mut nums: Vec<i32>, k: i32) -> Vec<i32> {
        let n = nums.len();
        let k = k as usize;
        // 因为返回字典序最小的，所以先reverse
        nums.reverse();
        // 下标从1开始
        let mut pre_sum = vec![0; n + 1];
        for i in 0..n {
            pre_sum[i + 1] = pre_sum[i] + nums[i];
        }

        let mut dp = vec![vec![0; 4]; n + 1];
        for i in k..=n {
            for j in 1..4 {
                dp[i][j] = dp[i - 1][j].max(dp[i - k][j - 1] + pre_sum[i] - pre_sum[i - k]);
            }
        }

        let mut res = vec![0; 3];
        let (mut i, mut j, mut index) = (n, 3, 0);
        while j > 0 {
            // 存在和更大的子数组，相比[i-k,i]
            if dp[i - 1][j] - dp[i - k][j - 1] > pre_sum[i] - pre_sum[i - k] {
                i -= 1;
            } else {
                res[index] = (n - i) as i32;
                index += 1;
                i -= k;
                j -= 1;
            }
        }
        res
    }
}
// @lc code=end
