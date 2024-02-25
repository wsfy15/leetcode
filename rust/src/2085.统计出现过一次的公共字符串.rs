/*
 * @lc app=leetcode.cn id=2085 lang=rust
 *
 * [2085] 统计出现过一次的公共字符串
 */

// @lc code=start
use std::collections::HashMap;

impl Solution {
    pub fn count_words(words1: Vec<String>, words2: Vec<String>) -> i32 {
        let mut m1 = HashMap::new();
        let mut m2 = HashMap::new();
        words1.iter().for_each(|s| {
            *m1.entry(s).or_insert(0) += 1;
        });
        words2.iter().for_each(|s| {
            *m2.entry(s).or_insert(0) += 1;
        });

        let mut res = 0;
        for (k, v) in m1 {
            if let Some(&u) = m2.get(k) {
                if v == 1 && u == 1 {
                    res += 1;
                }
            }
        }

        res
    }
}
// @lc code=end
