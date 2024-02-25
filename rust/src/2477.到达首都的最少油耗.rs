/*
 * @lc app=leetcode.cn id=2477 lang=rust
 *
 * [2477] 到达首都的最少油耗
 */

// @lc code=start
impl Solution {
    pub fn minimum_fuel_cost(roads: Vec<Vec<i32>>, seats: i32) -> i64 {
        let mut graph = vec![vec![]; roads.len() + 1];
        roads.iter().for_each(|road| {
            let (a, b) = (road[0], road[1]);
            graph[a as usize].push(b);
            graph[b as usize].push(a);
        });

        let mut res = 0;
        Self::dfs(&graph, seats, 0, 0, &mut res);
        res
    }

    fn dfs(graph: &Vec<Vec<i32>>, seats: i32, cur: usize, father: usize, res: &mut i64) -> i32 {
        let mut size = 1;
        for &n in &graph[cur] {
            if n as usize == father {
                continue;
            }

            size += Self::dfs(graph, seats, n as usize, cur, res);
        }
        if cur != 0 {
            *res += ((size + seats - 1) / seats) as i64;
        }
        size
    }
}
// @lc code=end
