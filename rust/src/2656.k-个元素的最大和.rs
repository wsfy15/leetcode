/*
 * @lc app=leetcode.cn id=2656 lang=rust
 *
 * [2656] K 个元素的最大和
 */

// @lc code=start
impl Solution {
    pub fn maximize_sum(nums: Vec<i32>, k: i32) -> i32 {
        let num = *nums.iter().max().unwrap();
        (num + num + k - 1) * k / 2
    }
}
// @lc code=end
