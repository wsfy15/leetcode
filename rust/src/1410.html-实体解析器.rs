/*
 * @lc app=leetcode.cn id=1410 lang=rust
 *
 * [1410] HTML 实体解析器
 */

// @lc code=start
impl Solution {
    pub fn entity_parser(text: String) -> String {
        let text = text.into_bytes();
        let n = text.len();
        let tab = [
            ("&quot;".to_string().into_bytes(), b'"'),
            ("&apos;".to_string().into_bytes(), b'\''),
            ("&amp;".to_string().into_bytes(), b'&'),
            ("&gt;".to_string().into_bytes(), b'>'),
            ("&lt;".to_string().into_bytes(), b'<'),
            ("&frasl;".to_string().into_bytes(), b'/'),
        ];

        let mut i = 0;
        let mut res = vec![];
        while i < n {
            if text[i] == b'&' {
                let mut ok = false;
                for t in tab.iter() {
                    let len = t.0.len();
                    if i + len > n {
                        continue;
                    }
                    if text[i..i + len] == t.0 {
                        res.push(t.1);
                        i += len;
                        ok = true;
                        break;
                    }
                }
                if ok {
                    continue;
                }
                res.push(b'&');
            } else {
                res.push(text[i]);
            }
            i += 1;
        }

        unsafe { String::from_utf8_unchecked(res) }
    }
}
// @lc code=end
