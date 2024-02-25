/*
 * @lc app=leetcode.cn id=2859 lang=rust
 *
 * [2859] 计算 K 置位下标对应元素的和
 */

// @lc code=start
impl Solution {
    pub fn sum_indices_with_k_set_bits(nums: Vec<i32>, k: i32) -> i32 {
        let mut res = 0;
        nums.iter().enumerate().for_each(|(i, &num)| {
            if i.count_ones() == k as u32 {
                res += num;
            }
        });

        res
    }
}
// @lc code=end
