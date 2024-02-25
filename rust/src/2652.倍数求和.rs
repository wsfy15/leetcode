/*
 * @lc app=leetcode.cn id=2652 lang=rust
 *
 * [2652] 倍数求和
 */

// @lc code=start
impl Solution {
    pub fn sum_of_multiples(n: i32) -> i32 {
        let mut res = 0;
        for i in 1..n + 1 {
            if i % 3 == 0 || i % 5 == 0 || i % 7 == 0 {
                res += i;
            }
        }

        res
    }
}
// @lc code=end
