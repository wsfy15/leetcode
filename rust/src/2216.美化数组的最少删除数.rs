/*
 * @lc app=leetcode.cn id=2216 lang=rust
 *
 * [2216] 美化数组的最少删除数
 */

// @lc code=start
impl Solution {
    pub fn min_deletion(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut res = 0;
        let mut i = 0;
        while i < n - 1 {
            if nums[i] == nums[i + 1] {
                res += 1;
            } else {
                i += 1;
            }
            i += 1;
        }

        if (n - res as usize) % 2 == 1 {
            res += 1;
        }

        res
    }
}
// @lc code=end
