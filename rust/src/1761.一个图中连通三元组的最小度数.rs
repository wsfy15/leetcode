/*
 * @lc app=leetcode.cn id=1761 lang=rust
 *
 * [1761] 一个图中连通三元组的最小度数
 */

// @lc code=start
impl Solution {
    // 对于一个连通三元组，总度数为三个点的总度数减去它们内部的度数(2*3=6)
    pub fn min_trio_degree(n: i32, edges: Vec<Vec<i32>>) -> i32 {
        let n = n as usize;
        let mut adj = vec![vec![false; n]; n];
        let mut deg = vec![0; n];
        edges.iter().for_each(|edge| {
            let (a, b) = (edge[0] as usize - 1, edge[1] as usize - 1);
            deg[a] += 1;
            deg[b] += 1;
            adj[a][b] = true;
            adj[b][a] = true;
        });

        let mut res = i32::MAX;
        for i in 0..n {
            for j in i + 1..n {
                if !adj[i][j] {
                    continue;
                }

                for k in j + 1..n {
                    if adj[i][k] && adj[j][k] {
                        res = res.min(deg[i] + deg[j] + deg[k] - 6);
                    }
                }
            }
        }

        if res == i32::MAX {
            -1
        } else {
            res
        }
    }
}
// @lc code=end
