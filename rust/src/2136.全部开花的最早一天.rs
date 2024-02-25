/*
 * @lc app=leetcode.cn id=2136 lang=rust
 *
 * [2136] 全部开花的最早一天
 */

// @lc code=start
impl Solution {
    pub fn earliest_full_bloom(plant_time: Vec<i32>, grow_time: Vec<i32>) -> i32 {
        let mut id: Vec<_> = (0..grow_time.len()).collect();
        let mut ans = 0;

        id.sort_unstable_by(|&i, &j| grow_time[j].cmp(&grow_time[i]));

        let mut plant_days = 0;
        for i in id {
            plant_days += plant_time[i];
            ans = ans.max(plant_days + grow_time[i]);
        }

        ans
    }
}
// @lc code=end
