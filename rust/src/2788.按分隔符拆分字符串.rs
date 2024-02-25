/*
 * @lc app=leetcode.cn id=2788 lang=rust
 *
 * [2788] 按分隔符拆分字符串
 */

// @lc code=start
impl Solution {
    pub fn split_words_by_separator(words: Vec<String>, separator: char) -> Vec<String> {
        let mut res = vec![];
        words.iter().for_each(|word| {
            for item in word.split(separator) {
                if !item.is_empty() {
                    res.push(item.to_string());
                }
            }
        });

        res
    }
}
// @lc code=end
