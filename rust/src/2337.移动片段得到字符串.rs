/*
 * @lc app=leetcode.cn id=2337 lang=rust
 *
 * [2337] 移动片段得到字符串
 */

// @lc code=start
impl Solution {
    pub fn can_change(start: String, target: String) -> bool {
        // L R 不能互相穿过
        let s1: String = start.chars().into_iter().filter(|&c| c != '_').collect();
        let t1: String = target.chars().into_iter().filter(|&c| c != '_').collect();
        if s1 != t1 {
            return false
        }
    
        let mut t_index = 0;
        let (sc, tc): (Vec<_>, Vec<_>) = (start.chars().collect(), target.chars().collect());
        let n = start.len();
        for i in 0..n {
            if sc[i] == '_' {
                continue;
            }
            while t_index < n && tc[t_index] == '_' {
                t_index += 1;
            }
    
            if i != t_index && ((sc[i] == 'L') == (i < t_index)) {
                return false;
            }
            t_index += 1;
        }
    
        true
    }
}
// @lc code=end

