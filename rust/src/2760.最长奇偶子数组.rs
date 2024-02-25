/*
 * @lc app=leetcode.cn id=2760 lang=rust
 *
 * [2760] 最长奇偶子数组
 */

// @lc code=start
impl Solution {
    pub fn longest_alternating_subarray(nums: Vec<i32>, threshold: i32) -> i32 {
        let mut res = 0;
        let mut start = 0;
        let n = nums.len();

        while start < n {
            if nums[start] % 2 == 1 || nums[start] > threshold {
                start += 1;
                continue;
            }

            let mut index = start + 1;
            while index < n {
                if nums[index] > threshold || (nums[index] % 2 == nums[index - 1] % 2) {
                    break;
                }
                index += 1;
            }
            res = res.max(index - start);
            start = index;
        }

        res as i32
    }
}
// @lc code=end
