/*
 * @lc app=leetcode.cn id=1944 lang=rust
 *
 * [1944] 队列中可以看到的人数
 */

// @lc code=start
impl Solution {
    pub fn can_see_persons_count(heights: Vec<i32>) -> Vec<i32> {
        let n = heights.len();
        let mut stack = vec![];
        let mut res = vec![0; n];
        for i in (0..n).rev() {
            while stack.len() > 0 && heights[i] > *stack.last().unwrap() {
                res[i] += 1;
                stack.pop();
            }

            if stack.len() > 0 {
                res[i] += 1;
            }
            stack.push(heights[i]);
        }

        res
    }
}
// @lc code=end
