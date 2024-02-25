/*
 * @lc app=leetcode.cn id=292 lang=rust
 *
 * [292] Nim 游戏
 */

// @lc code=start
impl Solution {
    pub fn can_win_nim(n: i32) -> bool {
        n & 3 > 0
    }
}
// @lc code=end
