/*
 * @lc app=leetcode.cn id=2707 lang=rust
 *
 * [2707] 字符串中的额外字符
 */

// @lc code=start
use std::collections::HashSet;

impl Solution {
    pub fn min_extra_char(s: String, dictionary: Vec<String>) -> i32 {
        let m = dictionary.into_iter().collect::<HashSet<_>>();
        let n = s.len();
        let mut dp = vec![-1; n];

        fn dfs(i: i32, dp: &mut [i32], m: &HashSet<String>, s: &str) -> i32 {
            if i < 0 {
                return 0;
            }

            if dp[i as usize] != -1 {
                return dp[i as usize];
            }
            let mut res = dfs(i - 1, dp, m, s) + 1;

            for j in 0..=i {
                if m.contains(&s[j as usize..=(i as usize)]) {
                    res = res.min(dfs(j - 1, dp, m, s));
                }
            }

            dp[i as usize] = res;
            res
        }

        dfs(n as i32 - 1, &mut dp, &m, &s)
    }
}
// @lc code=end
