/*
 * @lc app=leetcode.cn id=2609 lang=rust
 *
 * [2609] 最长平衡子字符串
 */

// @lc code=start
impl Solution {
    pub fn find_the_longest_balanced_substring(s: String) -> i32 {
        let mut res = 0;
        let n = s.len();
        let s: Vec<char> = s.chars().collect();
        let (mut zeros, mut ones) = (0, 0);
        for i in 0..n {
            if s[i] == '0' {
                if ones == 0 {
                    zeros += 1;
                } else {
                    ones = 0;
                    zeros = 1;
                }
            } else {
                ones += 1;
                if ones <= zeros {
                    res = res.max(ones * 2);
                }
            }
        }

        res
    }
}
// @lc code=end
