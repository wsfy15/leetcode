/*
 * @lc app=leetcode.cn id=447 lang=rust
 *
 * [447] 回旋镖的数量
 */

// @lc code=start
use std::collections::HashMap;

impl Solution {
    pub fn number_of_boomerangs(points: Vec<Vec<i32>>) -> i32 {
        let mut res = 0;
        let n = points.len();
        points.iter().enumerate().for_each(|(index, point)| {
            let mut counter: HashMap<i32, i32> = HashMap::new();
            for i in 0..n {
                if i == index {
                    continue;
                }

                let (xdiff, ydiff) = (points[i][0] - point[0], points[i][1] - point[1]);
                let distance = xdiff * xdiff + ydiff * ydiff;
                let c = *counter.get(&distance).unwrap_or(&0);
                res += 2 * c;
                counter.insert(distance, c + 1);
            }
        });

        res
    }
}
// @lc code=end
