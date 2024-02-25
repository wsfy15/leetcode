/*
 * @lc app=leetcode.cn id=2582 lang=rust
 *
 * [2582] 递枕头
 */

// @lc code=start
impl Solution {
    pub fn pass_the_pillow(n: i32, time: i32) -> i32 {
        let round = time / (n - 1);
        let offset = time % (n - 1);
        if round % 2 == 0 {
            return 1 + offset;
        }
        return n - offset;
    }
}
// @lc code=end
