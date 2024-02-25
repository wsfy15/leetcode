/*
 * @lc app=leetcode.cn id=1657 lang=rust
 *
 * [1657] 确定两个字符串是否接近
 */

// @lc code=start
impl Solution {
    pub fn close_strings(word1: String, word2: String) -> bool {
        // 没有不同的字符，字符出现的频次相同
        let (mut chars1, mut chars2) = (vec![0; 26], vec![0; 26]);
        word1.chars().for_each(|c| {
            chars1[c as usize - 'a' as usize] += 1;
        });
        word2.chars().for_each(|c| {
            chars2[c as usize - 'a' as usize] += 1;
        });

        for i in 0..26 {
            if chars1[i] + chars2[i] == 0 {
                continue;
            }
            if chars1[i] == 0 || chars2[i] == 0 {
                return false;
            }
        }

        chars1.sort_unstable();
        chars2.sort_unstable();
        for i in 0..26 {
            if chars1[i] != chars2[i] {
                return false;
            }
        }

        true
    }
}
// @lc code=end
