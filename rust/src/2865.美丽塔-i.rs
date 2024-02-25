/*
 * @lc app=leetcode.cn id=2865 lang=rust
 *
 * [2865] 美丽塔 I
 */

// @lc code=start
impl Solution {
    pub fn maximum_sum_of_heights(max_heights: Vec<i32>) -> i64 {
        let n = max_heights.len();
        let mut res = 0;
        let mut prefix = vec![0; n];
        let mut suffix = vec![0; n];
        let mut stack1 = vec![];
        let mut stack2 = vec![];

        for i in 0..n {
            while !stack1.is_empty() && max_heights[i] < max_heights[stack1[stack1.len() - 1]] {
                stack1.pop();
            }
            if stack1.is_empty() {
                prefix[i] = (i + 1) as i64 * max_heights[i] as i64;
            } else {
                prefix[i] = prefix[stack1[stack1.len() - 1]]
                    + (i - stack1[stack1.len() - 1]) as i64 * max_heights[i] as i64;
            }
            stack1.push(i);
        }

        for i in (0..n).rev() {
            while !stack2.is_empty() && max_heights[i] < max_heights[stack2[stack2.len() - 1]] {
                stack2.pop();
            }
            if stack2.is_empty() {
                suffix[i] = (n - i) as i64 * max_heights[i] as i64;
            } else {
                suffix[i] = suffix[stack2[stack2.len() - 1]]
                    + (stack2[stack2.len() - 1] - i) as i64 * max_heights[i] as i64;
            }
            stack2.push(i);
            res = res.max(prefix[i] + suffix[i] - max_heights[i] as i64);
        }

        res
    }
}
// @lc code=end
