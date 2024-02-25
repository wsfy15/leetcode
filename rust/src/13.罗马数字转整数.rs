/*
 * @lc app=leetcode.cn id=13 lang=rust
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
use std::collections::HashMap;

impl Solution {
    pub fn roman_to_int(s: String) -> i32 {
        let chars: Vec<char> = s.chars().collect();
        let mut m = HashMap::new();
        m.insert('I', 1);
        m.insert('V', 5);
        m.insert('X', 10);
        m.insert('L', 50);
        m.insert('C', 100);
        m.insert('D', 500);
        m.insert('M', 1000);
    
        let mut res = 0;
        let mut index = 0;
        let n = s.len();
        while index < n {
            let cur = m.get(&chars[index]).unwrap_or(&0);
            if index + 1 < n && cur < m.get(&chars[index+1]).unwrap_or(&0) {
                res += m.get(&chars[index+1]).unwrap() - cur;
                index += 1;
            } else {
                res += cur;
            }
            index += 1;
        }
    
        return res
    }
}
// @lc code=end

