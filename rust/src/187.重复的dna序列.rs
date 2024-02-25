/*
 * @lc app=leetcode.cn id=187 lang=rust
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
use std::collections::HashMap;

impl Solution {
    pub fn find_repeated_dna_sequences(s: String) -> Vec<String> {
        let s: Vec<char> = s.chars().collect();
        let mut map = HashMap::new();
        for w in s.windows(10) {
            *map.entry(w).or_insert(0) += 1;
        }

        map.into_iter()
            .filter(|&(_, cnt)| cnt > 1)
            .map(|(s, _)| s.iter().collect::<String>())
            .collect()
    }
}
// @lc code=end
