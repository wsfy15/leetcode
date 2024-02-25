/*
 * @lc app=leetcode.cn id=9 lang=rust
 *
 * [9] 回文数
 */

// @lc code=start
impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        if x < 0 || (x % 10 == 0 && x > 0) {
            return false;
        }
    
        let mut left = x;
        let mut right = 0;
        while left > right {
            right = right * 10 + left % 10;
            left /= 10;
        }
    
        // 对于奇数长度的数字 num/10 可以去掉最中间的数字
        return left == right || left == right/10
    }
}
// @lc code=end

