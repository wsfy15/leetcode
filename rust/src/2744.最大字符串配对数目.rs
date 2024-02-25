/*
 * @lc app=leetcode.cn id=2744 lang=rust
 *
 * [2744] 最大字符串配对数目
 */

// @lc code=start
impl Solution {
    pub fn maximum_number_of_string_pairs(words: Vec<String>) -> i32 {
        let mut seen = vec![vec![false; 26]; 26];
        let mut res = 0;

        for s in words {
            let s = s.as_bytes();
            let a = (s[0] - b'a') as usize;
            let b = (s[1] - b'a') as usize;

            if seen[b][a] {
                // 不需要，因为 words 包含的字符串互不相同
                //  seen[b][a] = true;
                res += 1;
            } else {
                seen[a][b] = true;
            }
        }

        res
    }
}
// @lc code=end
