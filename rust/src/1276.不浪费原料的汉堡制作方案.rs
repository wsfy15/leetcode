/*
 * @lc app=leetcode.cn id=1276 lang=rust
 *
 * [1276] 不浪费原料的汉堡制作方案
 */

// @lc code=start
impl Solution {
    pub fn num_of_burgers(tomato_slices: i32, cheese_slices: i32) -> Vec<i32> {
        if cheese_slices * 2 > tomato_slices
            || cheese_slices * 4 < tomato_slices
            || tomato_slices % 2 == 1
        {
            return vec![];
        }

        let total_jumbo = (tomato_slices - 2 * cheese_slices) / 2;
        let total_small = cheese_slices - total_jumbo;
        if total_jumbo >= 0 && total_small >= 0 {
            return vec![total_jumbo, total_small];
        }
        vec![]
    }
}
// @lc code=end
