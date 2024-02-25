/*
 * @lc app=leetcode.cn id=2316 lang=rust
 *
 * [2316] 统计无向图中无法互相到达点对数
 */

// @lc code=start
impl Solution {
    pub fn count_pairs(n: i32, edges: Vec<Vec<i32>>) -> i64 {
        let mut graph = vec![vec![]; n as usize];
        edges.iter().for_each(|edge| {
            graph[edge[0] as usize].push(edge[1]);
            graph[edge[1] as usize].push(edge[0]);
        });

        fn dfs(index: usize, graph: &Vec<Vec<i32>>, visited: &mut Vec<bool>) -> i32 {
            visited[index] = true;
            let mut size = 1;
            for &x in &graph[index] {
                if !visited[x as usize] {
                    size += dfs(x as usize, graph, visited);
                }
            }

            size
        }

        let mut visited = vec![false; n as usize];
        let mut res = 0;
        let mut total = 0;
        for i in 0..n as usize {
            if !visited[i] {
                let cur = dfs(i, &graph, &mut visited);
                res += cur as i64 * total as i64;
                total += cur;
            }
        }

        res
    }
}
// @lc code=end
