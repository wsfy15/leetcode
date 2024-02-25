/*
 * @lc app=leetcode.cn id=2670 lang=rust
 *
 * [2670] 找出不同元素数目差数组
 */

// @lc code=start
use std::collections::HashSet;

impl Solution {
    pub fn distinct_difference_array(nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len();
        let mut m = HashSet::new();
        let mut suffix = vec![0; n + 1];
        for i in (0..n).rev() {
            m.insert(nums[i]);
            suffix[i] = m.len();
        }

        let mut m = HashSet::new();
        let mut res = vec![0; n];
        for i in 0..n {
            m.insert(nums[i]);
            res[i] = m.len() as i32 - suffix[i + 1] as i32;
        }
        res
    }
}
// @lc code=end
