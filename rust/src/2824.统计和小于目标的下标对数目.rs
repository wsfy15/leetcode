/*
 * @lc app=leetcode.cn id=2824 lang=rust
 *
 * [2824] 统计和小于目标的下标对数目
 */

// @lc code=start
impl Solution {
    pub fn count_pairs(nums: Vec<i32>, target: i32) -> i32 {
        let mut nums = nums;
        nums.sort_unstable();
        let (mut left, mut right) = (0, nums.len() - 1);
        let mut res = 0;
        while left < right {
            if nums[left] + nums[right] < target {
                res += (right - left) as i32;
                left += 1;
            } else {
                right -= 1;
            }
        }

        res
    }
}
// @lc code=end
