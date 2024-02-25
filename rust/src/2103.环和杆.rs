/*
 * @lc app=leetcode.cn id=2103 lang=rust
 *
 * [2103] 环和杆
 */

// @lc code=start
impl Solution {
    pub fn count_points(rings: String) -> i32 {
        let n = rings.len();
        let mut flags = vec![0; 10];
        let bytes = rings.as_bytes();

        for i in (0..n).step_by(2) {
            let pos = bytes[i + 1] as usize - '0' as usize;
            flags[pos] |= if bytes[i] == b'R' {
                1
            } else if bytes[i] == b'G' {
                2
            } else {
                4
            };
        }

        flags.into_iter().filter(|&flag| flag == 7).count() as i32
    }
}
// @lc code=end
