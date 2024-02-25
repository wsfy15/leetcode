/*
 * @lc app=leetcode.cn id=1465 lang=rust
 *
 * [1465] 切割后面积最大的蛋糕
 */

// @lc code=start
impl Solution {
    pub fn max_area(
        h: i32,
        w: i32,
        mut horizontal_cuts: Vec<i32>,
        mut vertical_cuts: Vec<i32>,
    ) -> i32 {
        const MOD: i64 = 1e9 as i64 + 7;
        horizontal_cuts.sort_unstable();
        vertical_cuts.sort_unstable();

        let (mut max_width, mut max_height) = (vertical_cuts[0], horizontal_cuts[0]);

        for i in 1..horizontal_cuts.len() {
            max_height = max_height.max(horizontal_cuts[i] - horizontal_cuts[i - 1]);
        }
        max_height = max_height.max(h - horizontal_cuts[horizontal_cuts.len() - 1]);

        for i in 1..vertical_cuts.len() {
            max_width = max_width.max(vertical_cuts[i] - vertical_cuts[i - 1]);
        }
        max_width = max_width.max(w - vertical_cuts[vertical_cuts.len() - 1]);

        ((max_height as i64 * max_width as i64) % MOD) as i32
    }
}
// @lc code=end
