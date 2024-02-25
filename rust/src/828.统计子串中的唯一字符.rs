/*
 * @lc app=leetcode.cn id=828 lang=rust
 *
 * [828] 统计子串中的唯一字符
 */

// @lc code=start
impl Solution {
    pub fn unique_letter_string(s: String) -> i32 {
        let n = s.len();
        let mut indexes = vec![vec![-1; 1]; 26];
        for (i, c) in s.chars().enumerate() {
            indexes[c as usize - 'A' as usize].push(i as i32);
        }

        let mut res = 0;
        for v in indexes.iter_mut() {
            v.push(n as i32);
            for i in 1..v.len() - 1 {
                res += (v[i] - v[i - 1]) * (v[i + 1] - v[i]);
            }
        }
        res
    }
}
// @lc code=end
