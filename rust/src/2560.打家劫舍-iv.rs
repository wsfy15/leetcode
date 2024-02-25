/*
 * @lc app=leetcode.cn id=2560 lang=rust
 *
 * [2560] 打家劫舍 IV
 */

// @lc code=start
impl Solution {
    pub fn min_capability(nums: Vec<i32>, k: i32) -> i32 {
        let (mut left, mut right) = (
            nums.iter().min().unwrap().to_owned(),
            nums.iter().max().unwrap().to_owned(),
        );
        while left < right {
            let mid = left + (right - left) / 2;
            let mut count = 0;
            let mut i = 0;
            while i < nums.len() {
                if nums[i] <= mid {
                    count += 1;
                    i += 1;
                }
                i += 1;
            }

            if count >= k {
                right = mid;
            } else {
                left = mid + 1;
            }
        }

        left
    }
}
// @lc code=end
