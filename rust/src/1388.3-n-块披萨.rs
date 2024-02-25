/*
 * @lc app=leetcode.cn id=1388 lang=rust
 *
 * [1388] 3n 块披萨
 */

// @lc code=start
impl Solution {
    pub fn max_size_slices(slices: Vec<i32>) -> i32 {
        fn calc(s: &[i32]) -> i32 {
            let mut pre = vec![0; s.len()+2];
            for _ in 0..s.len()/3+1{
                let mut cur = vec![0; s.len()+2];
                for i in 2..pre.len() {
                    cur[i] = cur[i-1].max(pre[i-2]+s[i-2])
                }
                pre = cur; 
            }
            pre[s.len() + 1]
        }

        return calc(&slices[0..slices.len()-1]).max(calc(&slices[1..]))
    }
}
// @lc code=end

