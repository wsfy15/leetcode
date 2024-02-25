/*
 * @lc app=leetcode.cn id=137 lang=rust
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
impl Solution {
    pub fn single_number(nums: Vec<i32>) -> i32 {
        let (mut high, mut low) = (0, 0);
        nums.iter().for_each(|num| {
            low = low ^ num & (!high);
            high = high ^ num & (!low);
        });

        low
    }
}
// @lc code=end
