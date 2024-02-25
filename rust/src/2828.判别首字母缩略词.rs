/*
 * @lc app=leetcode.cn id=2828 lang=rust
 *
 * [2828] 判别首字母缩略词
 */

// @lc code=start
impl Solution {
    pub fn is_acronym(words: Vec<String>, s: String) -> bool {
        if words.len() != s.len() {
            return false;
        }

        let n = words.len();
        let bytes = s.as_bytes();
        for i in 0..n {
            if words[i].as_bytes()[0] != bytes[i] {
                return false;
            }
        }

        true
    }
}
// @lc code=end
