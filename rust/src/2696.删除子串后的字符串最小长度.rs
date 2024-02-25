/*
 * @lc app=leetcode.cn id=2696 lang=rust
 *
 * [2696] 删除子串后的字符串最小长度
 */

// @lc code=start
impl Solution {
    pub fn min_length(s: String) -> i32 {
        // let s = s.as_bytes();
        let mut stack = vec![];
        s.as_bytes().iter().for_each(|&c| {
            if stack.len() > 0
                && ((c == 'B' as u8 && *stack.last().unwrap() == 'A' as u8)
                    || (c == 'D' as u8 && *stack.last().unwrap() == 'C' as u8))
            {
                stack.pop();
            } else {
                stack.push(c);
            }
        });

        stack.len() as i32
    }
}
// @lc code=end
