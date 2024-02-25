/*
 * @lc app=leetcode.cn id=1488 lang=rust
 *
 * [1488] 避免洪水泛滥
 */

// @lc code=start
use std::collections::HashMap;

impl Solution {
    pub fn avoid_flood(rains: Vec<i32>) -> Vec<i32> {
        let mut res = vec![0; rains.len()];
        let mut rain_time = HashMap::new();
        for (index, pool) in rains.iter().enumerate() {
            if *pool > 0 {
                res[index] = -1;
                if let Some(&last_time) = rain_time.get(pool) {
                    let mut j = last_time + 1;
                    while j < index {
                        if res[j] == 0 {
                            res[j] = *pool;
                            break;
                        }
                        j += 1;
                    }
                    if j == index {
                        return vec![];
                    }
                }
            }

            rain_time.insert(*pool, index);
        }

        for i in 0..res.len() {
            if res[i] == 0 {
                res[i] = 1;
            }
        }

        res
    }
}
// @lc code=end
