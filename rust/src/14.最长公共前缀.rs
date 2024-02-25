/*
 * @lc app=leetcode.cn id=14 lang=rust
 *
 * [14] 最长公共前缀
 */

// @lc code=start
impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        let n = strs.len();
        if n == 0 {
            return "".to_string();
        }
    
        let strs: Vec<_> = strs.into_iter().map(|s| s.into_bytes()).collect();
        let mut index = 0;
        loop {
            if strs[0].len() <= index {
                return std::str::from_utf8(&strs[0][0..index]).unwrap().to_string();
            }
    
            for i in 1..n {
                if strs[i].len() <= index || strs[i][index] != strs[i-1][index] {
                    return std::str::from_utf8(&strs[0][0..index]).unwrap().to_string();
                }
            }
            index += 1;
        }
    }
}
// @lc code=end

