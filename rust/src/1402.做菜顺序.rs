/*
 * @lc app=leetcode.cn id=1402 lang=rust
 *
 * [1402] 做菜顺序
 */

// @lc code=start
impl Solution {
    pub fn max_satisfaction(mut satisfaction: Vec<i32>) -> i32 {
        satisfaction.sort_by(|a, b| b.cmp(a));

        let mut sum = 0;
        let mut res = 0;
        for i in 0..satisfaction.len() {
            sum += satisfaction[i];
            if sum < 0 {
                break;
            }

            res += sum;
        }

        res
    }
}
// @lc code=end
