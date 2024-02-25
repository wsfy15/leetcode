/*
 * @lc app=leetcode.cn id=2048 lang=rust
 *
 * [2048] 下一个更大的数值平衡数
 */

// @lc code=start
impl Solution {
    pub fn next_beautiful_number(n: i32) -> i32 {
        fn is_beautiful_number(n: i32) -> bool {
            let mut count = vec![0; 8];
            let mut n = n;
            while n > 0 {
                let digit = n % 10;
                if digit == 0 || digit > 6 {
                    return false;
                }
                n /= 10;
                count[digit as usize] += 1;
            }
            for i in 1..8 {
                if count[i as usize] > 0 && count[i as usize] != i {
                    return false;
                }
            }

            true
        }

        if n == 0 {
            return 1;
        }
        let mut n = n + 1;
        while !is_beautiful_number(n) {
            n += 1;
        }
        n
    }
}
// @lc code=end
