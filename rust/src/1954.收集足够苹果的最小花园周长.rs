/*
 * @lc app=leetcode.cn id=1954 lang=rust
 *
 * [1954] 收集足够苹果的最小花园周长
 */

// @lc code=start
impl Solution {
    pub fn minimum_perimeter(needed_apples: i64) -> i64 {
        let (mut left, mut right) = (0, 100000i64);
        while left < right {
            let mid = left + (right - left) / 2;
            if mid * (mid + 1) * (2 * mid + 1) * 2 < needed_apples {
                left = mid + 1;
            } else {
                right = mid;
            }
        }
        left * 8
    }
}
// @lc code=end
