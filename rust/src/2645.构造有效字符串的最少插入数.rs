/*
 * @lc app=leetcode.cn id=2645 lang=rust
 *
 * [2645] 构造有效字符串的最少插入数
 */

// @lc code=start
impl Solution {
    pub fn add_minimum(word: String) -> i32 {
        let s = word.as_bytes();
        // s[0] - 'a'   +   'c' - s[n-1]
        let mut ans = s[0] as i32 - *s.last().unwrap() as i32 + 2;

        for i in 1..s.len() {
            ans += (s[i] as i32 - s[i - 1] as i32 + 2) % 3;
        }
        ans
    }
}
// @lc code=end
