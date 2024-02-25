/*
 * @lc app=leetcode.cn id=2520 lang=rust
 *
 * [2520] 统计能整除数字的位数
 */

// @lc code=start
impl Solution {
    pub fn count_digits(num: i32) -> i32 {
        let digits: Vec<_> = num.to_string().bytes().collect();
        let mut res = 0;
        for digit in digits {
            if num % (digit & 0xf) as i32 == 0 {
                res += 1;
            }
        }

        res
    }
}
// @lc code=end
