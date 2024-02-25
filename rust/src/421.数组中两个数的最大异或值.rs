/*
 * @lc app=leetcode.cn id=421 lang=rust
 *
 * [421] 数组中两个数的最大异或值
 */

// @lc code=start
use std::collections::HashSet;

impl Solution {
    pub fn find_maximum_xor(nums: Vec<i32>) -> i32 {
        let mut res = 0;
        let mut mask = 0;
        let mx = *nums.iter().max().unwrap();
        let high_bit = 31 - mx.leading_zeros() as i32;

        let mut seen = HashSet::new();

        for i in (0..=high_bit).rev() {
            seen.clear();
            mask |= 1 << i;
            let new_res = res | (1 << i);
            for &x in &nums {
                let x = x & mask;
                if seen.contains(&(x ^ new_res)) {
                    res = new_res;
                    break;
                }
                seen.insert(x);
            }
        }

        res
    }
}
// @lc code=end
