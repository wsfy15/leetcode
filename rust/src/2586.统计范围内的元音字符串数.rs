/*
 * @lc app=leetcode.cn id=2586 lang=rust
 *
 * [2586] 统计范围内的元音字符串数
 */

// @lc code=start
impl Solution {
    pub fn vowel_strings(words: Vec<String>, left: i32, right: i32) -> i32 {
        let mut res = 0;
        for i in left..=right {
            if Self::is_vowel(&words[i as usize]) {
                res += 1
            }
        }

        res
    }

    fn is_vowel(word: &String) -> bool {
        return "aeiou".contains(word.chars().next().unwrap())
            && "aeiou".contains(word.chars().last().unwrap());
    }
}
// @lc code=end
