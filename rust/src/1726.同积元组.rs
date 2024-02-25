/*
 * @lc app=leetcode.cn id=1726 lang=rust
 *
 * [1726] 同积元组
 */

// @lc code=start
use std::collections::HashMap;

impl Solution {
    pub fn tuple_same_product(nums: Vec<i32>) -> i32 {
        let mut product_count = HashMap::new();
        for i in 0..nums.len() {
            for j in i + 1..nums.len() {
                *product_count.entry(nums[i] * nums[j]).or_insert(0) += 1;
            }
        }

        product_count
            .values()
            .map(|count| count * (count - 1) / 2 * 8)
            .sum()
    }
}
// @lc code=end
