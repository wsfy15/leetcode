/*
 * @lc app=leetcode.cn id=833 lang=rust
 *
 * [833] 字符串中的查找与替换
 */

// @lc code=start
impl Solution {
    pub fn find_replace_string(mut s: String, indices: Vec<i32>, sources: Vec<String>, targets: Vec<String>) -> String {
        let n = indices.len();
        let mut arr = vec![];
        for i in 0..n {
            arr.push((indices[i] as usize, i));
        }

        arr.sort_unstable();

        for (i, j) in arr.into_iter().rev() {
            let source = &sources[j];
            let m = source.len();
            let target = &targets[j];

            if i + m <= s.len() && &s[i..i+m] == source {
                s.replace_range(i..i+m, &target)
            }
        }
        s
    }
}
// @lc code=end

