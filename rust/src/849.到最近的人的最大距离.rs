/*
 * @lc app=leetcode.cn id=849 lang=rust
 *
 * [849] 到最近的人的最大距离
 */

// @lc code=start
impl Solution {
    pub fn max_dist_to_closest(seats: Vec<i32>) -> i32 {
        let n = seats.len();
        let mut last_seated_index = -1;
        let mut res = 1;
        for i in 0..n {
            if seats[i] == 1 {
                if last_seated_index == -1 {
                    res = res.max(i as i32);
                } else {
                    res = res.max((i as i32 - last_seated_index)/2);
                }
                last_seated_index = i as i32;
            }
        }
        res = res.max(n as i32 - last_seated_index - 1);
    
        return res;
    }
}
// @lc code=end

