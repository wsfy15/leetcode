/*
 * @lc app=leetcode.cn id=53 lang=rust
 *
 * [53] 最大子数组和
 */

// @lc code=start
impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        let (mut sum, mut res) = (0, nums[0]);

        nums.iter().for_each(|&num| {
            sum += num;
            if sum > res {
                res = sum;
            }
            if sum < 0 {
                sum = 0;
            }
        });

        res
    }
}
// @lc code=end
