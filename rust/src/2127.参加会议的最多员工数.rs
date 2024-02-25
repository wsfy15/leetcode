/*
 * @lc app=leetcode.cn id=2127 lang=rust
 *
 * [2127] 参加会议的最多员工数
 */

// @lc code=start
use std::collections::VecDeque;

impl Solution {
    pub fn maximum_invitations(favorite: Vec<i32>) -> i32 {
        let n = favorite.len();
        let mut deg = vec![0; n];
        for &f in &favorite {
            deg[f as usize] += 1; // 统计基环树每个节点的入度
        }

        let mut max_depth = vec![1; n];
        let mut q = VecDeque::new();
        for (i, &d) in deg.iter().enumerate() {
            if d == 0 {
                q.push_back(i);
            }
        }
        while let Some(x) = q.pop_front() {
            // 拓扑排序，剪掉图上所有树枝
            let y = favorite[x] as usize; // x 只有一条出边
            max_depth[y] = max_depth[x] + 1;
            deg[y] -= 1;
            if deg[y] == 0 {
                q.push_back(y);
            }
        }

        let mut max_ring_size = 0;
        let mut sum_chain_size = 0;
        for i in 0..n {
            if deg[i] == 0 {
                continue;
            }

            // 遍历基环上的点
            deg[i] = 0; // 将基环上的点的入度标记为 0，避免重复访问
            let mut ring_size = 1; // 基环长度
            let mut x = favorite[i] as usize;
            while x != i {
                deg[x] = 0; // 将基环上的点的入度标记为 0，避免重复访问
                ring_size += 1;
                x = favorite[x] as usize;
            }

            if ring_size == 2 {
                // 基环长度为 2
                sum_chain_size += max_depth[i] + max_depth[favorite[i] as usize];
            // 累加两条最长链的长度
            } else {
                max_ring_size = max_ring_size.max(ring_size); // 取所有基环长度的最大值
            }
        }
        max_ring_size.max(sum_chain_size)
    }
}
// @lc code=end
