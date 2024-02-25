/*
 * @lc app=leetcode.cn id=2808 lang=rust
 *
 * [2808] 使循环数组所有元素相等的最少秒数
 */

// @lc code=start
use std::collections::HashMap;
impl Solution {
    pub fn minimum_seconds(nums: Vec<i32>) -> i32 {
        let mut m: HashMap<i32, Vec<usize>> = HashMap::new();
        nums.iter().enumerate().for_each(|(i, &v)| {
            m.entry(v).or_insert(Vec::new()).push(i);
        });

        let n = nums.len();
        let mut res = n;
        for (_, mut val) in m {
            val.push(val[0] + n);
            let mut interval = 0;
            for i in 1..val.len() {
                interval = interval.max((val[i] - val[i - 1]) / 2);
            }

            res = res.min(interval);
        }

        res as i32
    }
}
// @lc code=end
