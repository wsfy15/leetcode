/*
 * @lc app=leetcode.cn id=162 lang=rust
 *
 * [162] 寻找峰值
 */

// @lc code=start
impl Solution {
    pub fn find_peak_element(nums: Vec<i32>) -> i32 {
        let (mut left, mut right) = (0, nums.len() - 1);
        while left < right {
            let mid = left + (right - left) / 2;
            // 右侧为升序，那么峰值肯定在右边，而且不可能是mid，所以left=mid+1
            if nums[mid] < nums[mid + 1] {
                left = mid + 1;
            } else {
                // 峰值在包含mid的左侧，有可能是mid
                right = mid;
            }
        }

        left as i32
    }
}
// @lc code=end
