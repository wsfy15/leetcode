/*
 * @lc app=leetcode.cn id=670 lang=rust
 *
 * [670] 最大交换
 */

// @lc code=start
impl Solution {
    pub fn maximum_swap(num: i32) -> i32 {
        let mut s = num.to_string().chars().collect::<Vec<char>>();
        let n = s.len();
        let mut max_index = n - 1;
        // p: 高位下标 q:低位下标
        let (mut p, mut q) = (n, 0);
        for i in (0..n - 1).rev() {
            if s[i] > s[max_index] {
                max_index = i;
            } else if s[i] < s[max_index] {
                p = i;
                q = max_index;
            }
        }

        if p == n {
            return num;
        }

        s.swap(p, q);
        s.iter().collect::<String>().parse::<i32>().unwrap()
    }
}
// @lc code=end
