/*
 * @lc app=leetcode.cn id=1654 lang=rust
 *
 * [1654] 到家的最少跳跃次数
 */

// @lc code=start
use std::collections::HashSet;

impl Solution {
    pub fn minimum_jumps(forbidden: Vec<i32>, a: i32, b: i32, x: i32) -> i32 {
        let mut set: HashSet<i32> = forbidden.into_iter().collect();
        // (position, count, can jump backward)
        let mut queue = vec![(0, 0, false)];

        while queue.len() > 0 {
            let (cur, count, used) = queue.remove(0);
            if cur == x {
                return count;
            }

            if cur + a < 6000 && !set.contains(&(cur + a)) {
                set.insert(cur + a);
                queue.push((cur + a, count + 1, false));
            }
            if cur - b > 0 && !used && !set.contains(&(cur - b)) {
                queue.push((cur - b, count + 1, true));
            }
        }
        -1
    }
}
// @lc code=end
