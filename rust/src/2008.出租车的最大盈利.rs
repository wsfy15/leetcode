/*
 * @lc app=leetcode.cn id=2008 lang=rust
 *
 * [2008] 出租车的最大盈利
 */

// @lc code=start
impl Solution {
    pub fn max_taxi_earnings(n: i32, rides: Vec<Vec<i32>>) -> i64 {
        let n = n as usize;
        let mut groups = vec![vec![]; n + 1];
        rides.iter().for_each(|ride| {
            let (start, end, tip) = (ride[0] as usize, ride[1] as usize, ride[2] as i64);
            groups[end].push((start, (end - start) as i64 + tip as i64));
        });

        let mut dp = vec![0i64; n + 1];
        for i in 2..=n {
            dp[i] = dp[i - 1];
            for &g in &groups[i] {
                dp[i] = dp[i].max(dp[g.0] + g.1);
            }
        }
        dp[n]
    }
}
// @lc code=end
