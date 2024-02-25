/*
 * @lc app=leetcode.cn id=2591 lang=rust
 *
 * [2591] 将钱分给最多的儿童
 */

// @lc code=start
impl Solution {
    pub fn dist_money(mut money: i32, mut children: i32) -> i32 {
        money -= children;
        if money < 0 {
            return -1;
        }

        let mut res = children.min(money / 7);
        money -= res * 7;
        children -= res;
        if (children == 0 && money > 0) || (children == 1 && money == 3) {
            res -= 1;
        }

        res
    }
}
// @lc code=end
