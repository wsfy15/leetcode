/*
 * @lc app=leetcode.cn id=1222 lang=rust
 *
 * [1222] 可以攻击国王的皇后
 */

// @lc code=start
use std::collections::HashSet;

impl Solution {
    pub fn queens_attackthe_king(queens: Vec<Vec<i32>>, king: Vec<i32>) -> Vec<Vec<i32>> {
        let mut positions = HashSet::new();
        queens.iter().for_each(|q| {
            positions.insert(q[0] * 8 + q[1]);
        });

        let mut res = vec![];
        for x in -1..2 {
            for y in -1..2 {
                if x == 0 && y == 0 {
                    continue;
                }

                let (mut row, mut col) = (king[0], king[1]);
                row += x;
                col += y;
                while row >= 0 && row < 8 && col >= 0 && col < 8 {
                    if positions.contains(&(row * 8 + col)) {
                        res.push(vec![row, col]);
                        break;
                    }
                    row += x;
                    col += y;
                }
            }
        }
        res
    }
}
// @lc code=end
