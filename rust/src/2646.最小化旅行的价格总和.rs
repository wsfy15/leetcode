/*
 * @lc app=leetcode.cn id=2646 lang=rust
 *
 * [2646] 最小化旅行的价格总和
 */

// @lc code=start
impl Solution {
    pub fn minimum_total_price(
        n: i32,
        edges: Vec<Vec<i32>>,
        price: Vec<i32>,
        trips: Vec<Vec<i32>>,
    ) -> i32 {
        let n = n as usize;
        let mut graph = vec![vec![]; n];
        edges.iter().for_each(|edge| {
            let (a, b) = (edge[0] as usize, edge[1] as usize);
            graph[a].push(b);
            graph[b].push(a);
        });

        fn dfs(
            graph: &Vec<Vec<usize>>,
            father: usize,
            cur: usize,
            end: usize,
            count: &mut Vec<i32>,
        ) -> bool {
            if cur == end {
                count[cur] += 1;
                return true;
            }
            for &next in &graph[cur] {
                if next != father && dfs(graph, cur, next, end, count) {
                    count[cur] += 1;
                    return true;
                }
            }

            false
        }

        let mut count = vec![0; n];
        trips.iter().for_each(|trip| {
            let (start, end) = (trip[0], trip[1]);
            dfs(
                &graph,
                start as usize,
                start as usize,
                end as usize,
                &mut count,
            );
        });

        fn dp(
            cur: usize,
            father: usize,
            graph: &Vec<Vec<usize>>,
            count: &Vec<i32>,
            price: &Vec<i32>,
        ) -> (i32, i32) {
            let mut not_half = count[cur] * price[cur];
            let mut half = not_half / 2;
            for &next in &graph[cur] {
                if next != father {
                    let (nh, h) = dp(next, cur, graph, count, price);
                    not_half += nh.min(h);
                    half += nh;
                }
            }

            (not_half, half)
        }

        let (nh, h) = dp(0, 0, &graph, &count, &price);

        nh.min(h)
    }
}
// @lc code=end
