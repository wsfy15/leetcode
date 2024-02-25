/*
 * @lc app=leetcode.cn id=2697 lang=rust
 *
 * [2697] 字典序最小回文串
 */

// @lc code=start
impl Solution {
    pub fn make_smallest_palindrome(s: String) -> String {
        let n = s.len();

        let mut s = s.into_bytes();
        // let mut s = s.as_bytes().to_vec(); //可保留s所有权
        // let mut s: Vec<char> = s.chars().collect();
        for i in 0..n / 2 {
            let j = n - i - 1;
            if s[i] != s[j] {
                if s[i] < s[j] {
                    s[j] = s[i];
                } else {
                    s[i] = s[j]
                }
            }
        }

        String::from_utf8(s).unwrap()
    }
}
// @lc code=end
