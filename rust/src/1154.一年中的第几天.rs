/*
 * @lc app=leetcode.cn id=1154 lang=rust
 *
 * [1154] 一年中的第几天
 */

// @lc code=start
impl Solution {
    pub fn day_of_year(date: String) -> i32 {
        let mut months = vec![0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
        let items = date
            .split("-")
            .map(|x| x.parse::<i32>().unwrap())
            .collect::<Vec<_>>();
        let year = items[0];
        if year % 4 == 0 && (year % 100 != 0 || year % 400 == 0) {
            months[2] += 1;
        }

        let mut res = items[2];
        for i in 1..items[1] as usize {
            res += months[i];
        }
        res
    }
}
// @lc code=end
