/*
 * @lc app=leetcode.cn id=1696 lang=rust
 *
 * [1696] 跳跃游戏 VI
 */

// @lc code=start
use std::collections::VecDeque;

impl Solution {
    pub fn max_result(nums: Vec<i32>, k: i32) -> i32 {
        let n = nums.len();
        let k = k as usize;
        let mut f = vec![0; n];
        f[0] = nums[0];
        let mut q = VecDeque::new();
        q.push_back(0);
        for i in 1..n {
            // 1. 出
            if *q.front().unwrap() + k < i {
                q.pop_front();
            }
            // 2. 转移
            f[i] = f[*q.front().unwrap()] + nums[i];
            // 3. 入
            while !q.is_empty() && f[i] >= f[*q.back().unwrap()] {
                q.pop_back();
            }
            q.push_back(i);
        }
        f[n - 1]
    }

    pub fn max_result2(nums: Vec<i32>, k: i32) -> i32 {
        let n = nums.len();
        let k = k as usize;
        let mut dp = vec![0; n];
        dp[0] = nums[0];
        for i in 1..n {
            dp[i] = dp[i.saturating_sub(k)..i].iter().max().unwrap() + nums[i];
        }
        dp[n - 1]
    }
}
// @lc code=end
