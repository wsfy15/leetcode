/*
 * @lc app=leetcode.cn id=1466 lang=rust
 *
 * [1466] 重新规划路线
 */

// @lc code=start
impl Solution {
    pub fn min_reorder(n: i32, connections: Vec<Vec<i32>>) -> i32 {
        let mut conn = vec![vec![]; n as usize];
        connections.iter().enumerate().for_each(|(i, c)| {
            let (a, b) = (c[0], c[1]);
            conn[a as usize].push(i);
            conn[b as usize].push(i);
        });

        let mut queue = vec![0];
        let mut res = 0;
        let mut visited = vec![false; connections.len()];

        while queue.len() > 0 {
            let cur = queue.remove(0);
            for &i in &conn[cur] {
                if visited[i] {
                    continue;
                }
                visited[i] = true;

                let (from, to) = (connections[i][0] as usize, connections[i][1] as usize);
                if from == cur {
                    res += 1;
                    queue.push(to);
                } else {
                    queue.push(from);
                }
            }
        }

        res
    }
}
// @lc code=end
