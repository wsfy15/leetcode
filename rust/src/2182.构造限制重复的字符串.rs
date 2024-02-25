/*
 * @lc app=leetcode.cn id=2182 lang=rust
 *
 * [2182] 构造限制重复的字符串
 */

// @lc code=start
impl Solution {
    pub fn repeat_limited_string(s: String, repeat_limit: i32) -> String {
        let mut counter = vec![0; 26];
        s.as_bytes().iter().for_each(|&c| {
            counter[(c - 'a' as u8) as usize] += 1;
        });

        let mut res = vec![];
        for i in (0..26 as usize).rev() {
            if counter[i] == 0 {
                continue;
            }

            let mut k = i as i32 - 1;
            loop {
                let mut j = 0;
                while j < repeat_limit && counter[i] > 0 {
                    res.push((i + 'a' as usize) as u8);
                    j += 1;
                    counter[i] -= 1;
                }

                // i 用完了，找下一个更小的字母
                if counter[i] == 0 {
                    break;
                }

                while k >= 0 && counter[k as usize] == 0 {
                    k -= 1;
                }
                if k < 0 {
                    break;
                }

                counter[k as usize] -= 1;
                res.push((k + 'a' as i32) as u8);
            }
        }

        String::from_utf8(res).unwrap()
    }
}
// @lc code=end
