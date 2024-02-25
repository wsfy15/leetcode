/*
 * @lc app=leetcode.cn id=318 lang=rust
 *
 * [318] 最大单词长度乘积
 */

// @lc code=start
impl Solution {
    pub fn max_product(words: Vec<String>) -> i32 {
        let n = words.len();
        let mut bits = vec![0; n];
        for i in 0..n {
            words[i]
                .bytes()
                .for_each(|b| bits[i] |= 1 << (b - 'a' as u8))
        }

        let mut res = 0;
        for i in 0..n {
            for j in i + 1..n {
                if bits[i] & bits[j] == 0 {
                    res = res.max(words[i].len() * words[j].len());
                }
            }
        }
        res as i32
    }
}
// @lc code=end
