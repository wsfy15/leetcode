/*
 * @lc app=leetcode.cn id=2846 lang=rust
 *
 * [2846] 边权重均等查询
 */

// @lc code=start
impl Solution {
    /// 倍增求 LCA
    pub fn min_operations_queries(
        n: i32,
        edges: Vec<Vec<i32>>,
        queries: Vec<Vec<i32>>,
    ) -> Vec<i32> {
        let n = n as usize;
        let mut g = vec![vec![]; n];
        for edge in edges {
            let (u, v, w) = (edge[0] as usize, edge[1] as usize, edge[2] - 1);
            g[u].push((v, w));
            g[v].push((u, w));
        }
        fn dfs(
            g: &Vec<Vec<(usize, i32)>>,
            u: usize,
            fa: usize,
            d: i32,
            cur: &mut [i32; 26],
            depth: &mut Vec<i32>,
            parent: &mut Vec<Vec<usize>>,
            info: &mut Vec<[i32; 26]>,
        ) {
            depth[u] = d;
            info[u] = cur.clone();
            for &(v, w) in &g[u] {
                if v != fa {
                    parent[0][v] = u;
                    cur[w as usize] += 1;
                    dfs(g, v, u, d + 1, cur, depth, parent, info);
                    cur[w as usize] -= 1;
                }
            }
        }
        let mx = 64 - n.leading_zeros() as usize;
        let mut parent = vec![vec![n; n]; mx]; // parent[i][j] 表示j节点向上跳 2^i 的节点。跳数放第一维更快（CPU缓存 ）
        let mut info = vec![[0; 26]; n];
        let mut depth = vec![0; n];
        dfs(
            &g,
            0,
            0,
            0,
            &mut [0; 26],
            &mut depth,
            &mut parent,
            &mut info,
        );
        for j in 0..mx - 1 {
            for i in 0..n {
                let pp = parent[j][i];
                if parent[j][i] != n {
                    parent[j + 1][i] = parent[j][pp];
                }
            }
        }
        queries
            .into_iter()
            .map(|query| {
                let (mut u, mut v) = (query[0] as usize, query[1] as usize);
                if depth[u] > depth[v] {
                    std::mem::swap(&mut u, &mut v);
                }
                let diff = depth[v] - depth[u];
                for i in 0..mx {
                    if diff >> i & 1 == 1 {
                        v = parent[i][v];
                    }
                }
                if u != v {
                    for i in (0..mx).rev() {
                        if parent[i][u] != parent[i][v] {
                            u = parent[i][u];
                            v = parent[i][v];
                        }
                    }
                    u = parent[0][u];
                }
                let lca = u;
                let (u, v) = (query[0] as usize, query[1] as usize);
                let (sum, max) = info[u]
                    .iter()
                    .zip(&info[v])
                    .zip(&info[lca])
                    .map(|((a, b), c)| *a + *b - *c * 2)
                    .fold((0, 0), |(sum, max), num| (sum + num, max.max(num)));
                sum - max
            })
            .collect()
    }
}
// @lc code=end
