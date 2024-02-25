/*
 * @lc app=leetcode.cn id=2731 lang=rust
 *
 * [2731] 移动机器人
 */

// @lc code=start
impl Solution {
    pub fn sum_distance(nums: Vec<i32>, s: String, d: i32) -> i32 {
        const MOD: i64 = 1_000_000_007;
        let mut a = vec![0; nums.len()];

        let s = s.as_bytes();
        for (index, &pos) in nums.iter().enumerate() {
            let d = if s[index] == 'L' as u8 { -d } else { d };
            a[index] = (pos + d) as i64;
        }
        a.sort();

        let mut res = 0;
        let mut sum = 0;
        for (index, &val) in a.iter().enumerate() {
            res = (res + index as i64 * val - sum) % MOD;
            sum += val;
        }

        res as i32
    }
}
// @lc code=end
