/*
 * @lc app=leetcode.cn id=2171 lang=rust
 *
 * [2171] 拿出最少数目的魔法豆
 */

// @lc code=start
impl Solution {
    pub fn minimum_removal(mut beans: Vec<i32>) -> i64 {
        beans.sort_unstable();
        let n = beans.len();
        let (mut keep, mut sum) = (0i64, 0i64);
        beans.iter().enumerate().for_each(|(i, &val)| {
            let val = val as i64;
            keep = keep.max((n - i) as i64 * val);
            sum += val;
        });
        sum - keep
    }
}
// @lc code=end
