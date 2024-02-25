/*
 * @lc app=leetcode.cn id=70 lang=rust
 *
 * [70] 爬楼梯
 */

// @lc code=start
impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
        if n < 3 {
            return n;
        }

        let (mut one, mut two) = (1, 2);
        (3..n).into_iter().for_each(|_| {
            let new_two = one + two;
            one = two;
            two = new_two;
        });
        one + two
    }
}
// @lc code=end
