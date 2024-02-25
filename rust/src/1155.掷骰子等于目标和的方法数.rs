/*
 * @lc app=leetcode.cn id=1155 lang=rust
 *
 * [1155] 掷骰子等于目标和的方法数
 */

// @lc code=start
impl Solution {
    pub fn num_rolls_to_target(n: i32, k: i32, target: i32) -> i32 {
        const MOD: i32 = 1_000_000_000 + 7;
        let n = n as usize;
        let k = k as usize;
        let target = target as usize;
        let mut dp = vec![vec![0; target + 1]; n + 1];
        dp[0][0] = 1;
        for i in 1..=n {
            for j in 0..=target {
                for m in 1..=k {
                    if j >= m {
                        dp[i][j] = (dp[i][j] + dp[i - 1][j - m]) % MOD;
                    }
                }
            }
        }

        dp[n][target]
    }
}
// @lc code=end
