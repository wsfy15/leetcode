/*
 * @lc app=leetcode.cn id=2603 lang=rust
 *
 * [2603] 收集树中金币
 */

// @lc code=start
impl Solution {
    pub fn collect_the_coins(coins: Vec<i32>, edges: Vec<Vec<i32>>) -> i32 {
        let n = coins.len();
        let mut left_edges = n as i32 - 1;
        let mut graph = vec![vec![]; n];
        let mut degree = vec![0; n];
        edges.iter().for_each(|edge| {
            let (a, b) = (edge[0] as usize, edge[1] as usize);
            degree[a] += 1;
            degree[b] += 1;
            graph[a].push(b);
            graph[b].push(a);
        });

        // 排除所有没有金币的叶子节点，以及排除掉后产生的新无金币叶子节点
        let mut queue = vec![];
        for i in 0..n {
            if degree[i] == 1 && coins[i] == 0 {
                queue.push(i);
            }
        }

        while queue.len() > 0 {
            let x = queue.remove(0);
            left_edges -= 1;
            for &n in &graph[x] {
                degree[n] -= 1;
                if degree[n] == 1 && coins[n] == 0 {
                    queue.push(n);
                }
            }
        }

        // 现在所有叶子节点都是有金币的，但因为可以 收集距离当前节点距离为 2 以内的所有金币
        // 这些叶子节点到父节点、父节点到其父节点的边也可以删除
        for i in 0..n {
            if degree[i] == 1 && coins[i] == 1 {
                queue.push(i);
            }
        }

        left_edges -= queue.len() as i32;
        for i in queue {
            for &y in &graph[i] {
                degree[y] -= 1;
                if degree[y] == 1 {
                    left_edges -= 1;
                }
            }
        }

        0.max(left_edges * 2) as i32
    }
}
// @lc code=end
