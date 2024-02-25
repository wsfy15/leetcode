/*
 * @lc app=leetcode.cn id=1921 lang=rust
 *
 * [1921] 消灭怪物的最大数量
 */

// @lc code=start
impl Solution {
    pub fn eliminate_maximum(dist: Vec<i32>, speed: Vec<i32>) -> i32 {
        let n = dist.len();
        let mut ttl: Vec<_> = (0..n)
            .map(|i| (dist[i] + speed[i] - 1) / speed[i])
            .collect();

        ttl.sort();
        for i in 0..n {
            if ttl[i] <= i as i32 {
                return i as i32;
            }
        }
        n as i32
    }
}
// @lc code=end
