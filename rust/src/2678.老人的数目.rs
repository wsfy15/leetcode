/*
 * @lc app=leetcode.cn id=2678 lang=rust
 *
 * [2678] 老人的数目
 */

// @lc code=start
impl Solution {
    pub fn count_seniors(details: Vec<String>) -> i32 {
        let mut res = 0;
        for detail in details {
            if detail[11..13].parse::<i32>().unwrap() > 60 {
                res += 1;
            }
        }

        res
    }
}
// @lc code=end
