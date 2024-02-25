/*
 * @lc app=leetcode.cn id=2605 lang=rust
 *
 * [2605] 从两个数字数组里生成最小数字
 */

// @lc code=start
use std::collections::HashSet;

impl Solution {
    pub fn min_number(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
        let mut nums1 = nums1;
        nums1.sort();
        let set: HashSet<_> = nums2.iter().collect();
        for i in &nums1 {
            if set.contains(i) {
                return *i;
            }
        }

        let min_num = nums2.iter().min().unwrap();
        (nums1[0] + min_num * 10).min(nums1[0] * 10 + min_num)
    }
}
// @lc code=end
