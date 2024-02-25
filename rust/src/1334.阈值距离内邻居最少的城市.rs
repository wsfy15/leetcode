/*
 * @lc app=leetcode.cn id=1334 lang=rust
 *
 * [1334] 阈值距离内邻居最少的城市
 */

// @lc code=start
impl Solution {
    pub fn find_the_city(n: i32, edges: Vec<Vec<i32>>, distance_threshold: i32) -> i32 {
        let n = n as usize;
        let (mut res, mut neighs) = (-1, std::i32::MAX >> 1);
        let mut g = vec![vec![std::i32::MAX >> 1; n]; n];
        edges.iter().for_each(|edge| {
            let (from, to, weight) = (edge[0], edge[1], edge[2]);
            g[from as usize][to as usize] = weight;
            g[to as usize][from as usize] = weight;
        });

        for k in 0..n {
            g[k][k] = 0;
            for i in 0..n {
                for j in 0..n {
                    g[i][j] = g[i][j].min(g[i][k] + g[k][j]);
                }
            }
        }

        for i in 0..n {
            let mut cnt = 0;
            for j in 0..n {
                if g[i][j] <= distance_threshold {
                    cnt += 1;
                }
            }
            if cnt <= neighs {
                res = i as i32;
                neighs = cnt;
            }
        }

        res
    }
}
// @lc code=end
