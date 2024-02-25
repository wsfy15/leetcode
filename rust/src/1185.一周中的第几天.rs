/*
 * @lc app=leetcode.cn id=1185 lang=rust
 *
 * [1185] 一周中的第几天
 */

// @lc code=start
impl Solution {
    pub fn day_of_the_week(day: i32, month: i32, year: i32) -> String {
        let (mut month, mut year) = (month, year);
        if month < 3 {
            month += 12;
            year -= 1;
        }

        let index =
            (day + 2 * month + 3 * (month + 1) / 5 + year + year / 4 - year / 100 + year / 400) % 7;
        let weekdays = vec![
            "Monday",
            "Tuesday",
            "Wednesday",
            "Thursday",
            "Friday",
            "Saturday",
            "Sunday",
        ];

        weekdays[index as usize].to_string()
    }
}
// @lc code=end
