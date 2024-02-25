/*
 * @lc app=leetcode.cn id=2511 lang=rust
 *
 * [2511] 最多可以摧毁的敌人城堡数目
 */

// @lc code=start
impl Solution {
    pub fn capture_forts(forts: Vec<i32>) -> i32 {
        let mut res = 0;
        let mut last_index: i32 = -1;
        let n = forts.len();
        for i in 0..n {
            if forts[i] != 0 {
                if last_index >= 0 && forts[last_index as usize] == -forts[i] {
                    res = res.max(i as i32 - last_index - 1);
                }
                last_index = i as i32;
            }
        }

        res
    }
}
// @lc code=end
