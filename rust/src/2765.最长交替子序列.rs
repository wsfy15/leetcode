/*
 * @lc app=leetcode.cn id=2765 lang=rust
 *
 * [2765] 最长交替子序列
 */

// @lc code=start
impl Solution {
    pub fn alternating_subarray(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut res = -1;
        let mut i = 0;
        while i < n - 1 {
            if nums[i + 1] - nums[i] != 1 {
                i += 1;
                continue;
            }

            let start = i;
            i += 2;
            while i < n && nums[i] == nums[start] + (i - start) as i32 % 2 {
                i += 1;
            }

            res = res.max((i - start) as i32);
            i -= 1;
        }

        res
    }
}
// @lc code=end
