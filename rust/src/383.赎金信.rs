/*
 * @lc app=leetcode.cn id=383 lang=rust
 *
 * [383] 赎金信
 */

// @lc code=start
impl Solution {
    pub fn can_construct(ransom_note: String, magazine: String) -> bool {
        let mut counter = vec![0; 26];
        magazine.as_bytes().iter().for_each(|&c| {
            counter[(c - 'a' as u8) as usize] += 1;
        });

        let bytes = ransom_note.as_bytes();
        for i in 0..bytes.len() {
            let index = (bytes[i] - 'a' as u8) as usize;
            counter[index] -= 1;
            if counter[index] < 0 {
                return false;
            }
        }
        true
    }
}
// @lc code=end
