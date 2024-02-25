/*
 * @lc app=leetcode.cn id=1 lang=rust
 *
 * [1] 两数之和
 */

// @lc code=start
use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        // num -> index
        let mut position = HashMap::with_capacity(nums.len());

        for i in 0..nums.len() {
            if let Some(k) = position.get(&(target - nums[i])) {
                return vec![*k as i32, i as i32];
            }
           
            position.insert(nums[i], i);
        }
        return vec![];
    }
}
// @lc code=end

