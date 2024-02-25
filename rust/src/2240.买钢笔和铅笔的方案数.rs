/*
 * @lc app=leetcode.cn id=2240 lang=rust
 *
 * [2240] 买钢笔和铅笔的方案数
 */

// @lc code=start
impl Solution {
    pub fn ways_to_buy_pens_pencils(total: i32, cost1: i32, cost2: i32) -> i64 {
        let max_count1 = total / cost1;
        let mut res: i64 = 0;
        (0..max_count1 + 1).for_each(|x| {
            res += ((total - x * cost1) / cost2 + 1) as i64;
        });
        res
    }
}
// @lc code=end
