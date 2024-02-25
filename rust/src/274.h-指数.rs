/*
 * @lc app=leetcode.cn id=274 lang=rust
 *
 * [274] H 指数
 */

// @lc code=start
impl Solution {
    pub fn h_index(mut citations: Vec<i32>) -> i32 {
        citations.sort();
        let n = citations.len() as i32;
        for i in 0..n {
            if citations[i as usize] >= n - i {
                return n - i;
            }
        }

        0
    }
}
// @lc code=end
