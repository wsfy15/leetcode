/*
 * @lc app=leetcode.cn id=2525 lang=rust
 *
 * [2525] 根据规则将箱子分类
 */

// @lc code=start
impl Solution {
    pub fn categorize_box(length: i32, width: i32, height: i32, mass: i32) -> String {
        let bulk = (length as i64) * (width as i64) * (height as i64);

        let bulky = bulk as f64 >= 1e9
            || length as f64 >= 1e4
            || width as f64 >= 1e4
            || height as f64 >= 1e4;
        let heavy = mass >= 100;

        if bulky && heavy {
            return "Both".to_string();
        }
        if bulky {
            return "Bulky".to_string();
        }
        if heavy {
            return "Heavy".to_string();
        }
        return "Neither".to_string();
    }
}
// @lc code=end
