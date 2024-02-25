/*
 * @lc app=leetcode.cn id=275 lang=rust
 *
 * [275] H 指数 II
 */

// @lc code=start
impl Solution {
    pub fn h_index(citations: Vec<i32>) -> i32 {
        let n = citations.len() as i32;
        if n == 0 || citations[n as usize - 1] == 0 {
            return 0;
        }

        let (mut left, mut right) = (0, n - 1);

        while left < right {
            let mid = left + (right - left) / 2;
            if citations[mid as usize] < n - mid {
                left = mid + 1;
            } else {
                right = mid;
            }
        }

        n - left
    }
}
// @lc code=end
