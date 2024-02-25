/*
 * @lc app=leetcode.cn id=2578 lang=rust
 *
 * [2578] 最小和分割
 */

// @lc code=start
impl Solution {
    pub fn split_num(num: i32) -> i32 {
        let mut s: Vec<u8> = num.to_string().bytes().collect();
        s.sort_unstable();
        let mut nums = [0, 0];
        for (index, &c) in s.iter().enumerate() {
            nums[index % 2] = nums[index % 2] * 10 + c as i32 - '0' as i32;
        }
        nums[0] + nums[1]
    }
}
// @lc code=end
