/*
 * @lc app=leetcode.cn id=228 lang=rust
 *
 * [228] 汇总区间
 */

// @lc code=start
impl Solution {
    pub fn summary_ranges(nums: Vec<i32>) -> Vec<String> {
        let n = nums.len();
        if n == 0 {
            return vec![];
        }

        let mut res = vec![];
        let mut start_index = 0;
        for i in 1..n {
            if nums[i] == nums[i - 1] + 1 {
                continue;
            }

            if i > start_index + 1 {
                res.push(format!("{}->{}", nums[start_index], nums[i - 1]));
            } else {
                res.push(format!("{}", nums[start_index]));
            }
            start_index = i;
        }

        if start_index < n - 1 {
            res.push(format!("{}->{}", nums[start_index], nums[n - 1]));
        } else {
            res.push(format!("{}", nums[start_index]));
        }

        res
    }
}
// @lc code=end
