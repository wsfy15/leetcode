/*
 * @lc app=leetcode.cn id=2682 lang=rust
 *
 * [2682] 找出转圈游戏输家
 */

// @lc code=start
use std::collections::HashSet;
impl Solution {
    pub fn circular_game_losers(n: i32, k: i32) -> Vec<i32> {
        let mut set = HashSet::new();
        let mut cur = 0;
        let mut round = 1;
        loop {
            set.insert(cur);
            cur = (cur + round * k) % n;
            if set.contains(&cur) {
                break;
            }
            round += 1;
        }
    
        let mut res = vec![];
        for i in 0..n {
            if !set.contains(&i) {
                res.push(i+1);
            }
        }
    
        return res
    }
}
// @lc code=end

