/*
 * @lc app=leetcode.cn id=2258 lang=rust
 *
 * [2258] 逃离火灾
 */

// @lc code=start
impl Solution {
    pub fn maximum_minutes(grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();
        let n = grid[0].len();
        // 返回三个数，分别表示到达安全屋/安全屋左边/安全屋上边的最短时间
        let bfs = |q: &mut Vec<(usize, usize)>| -> (i32, i32, i32) {
            let mut time = vec![vec![-1; n]; m]; // -1 表示未访问
            for &(i, j) in q.iter() {
                time[i][j] = 0;
            }
            let mut t = 1;
            while !q.is_empty() {
                // 每次循环向外扩展一圈
                let mut tmp = Vec::new();
                for &(i, j) in q.iter() {
                    for &(x, y) in &[(i - 1, j), (i + 1, j), (i, j - 1), (i, j + 1)] {
                        // 上下左右
                        if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 0 && time[x][y] < 0 {
                            time[x][y] = t;
                            tmp.push((x, y));
                        }
                    }
                }
                t += 1;
                *q = tmp;
            }
            (time[m - 1][n - 1], time[m - 1][n - 2], time[m - 2][n - 1])
        };

        let (man_to_house_time, m1, m2) = bfs(&mut vec![(0, 0)]);
        if man_to_house_time < 0 {
            // 人无法到安全屋
            return -1;
        }

        let mut fire_pos = Vec::new();
        for (i, row) in grid.iter().enumerate() {
            for (j, &x) in row.iter().enumerate() {
                if x == 1 {
                    fire_pos.push((i, j));
                }
            }
        }
        let (fire_to_house_time, f1, f2) = bfs(&mut fire_pos); // 多个着火点同时跑 BFS
        if fire_to_house_time < 0 {
            // 火无法到安全屋
            return 1_000_000_000;
        }

        let d = fire_to_house_time - man_to_house_time;
        if d < 0 {
            // 火比人先到安全屋
            return -1;
        }

        if m1 != -1 && m1 + d < f1 || // 安全屋左边相邻格子，人比火先到
           m2 != -1 && m2 + d < f2
        {
            // 安全屋上边相邻格子，人比火先到
            return d; // 图中第一种情况
        }
        return d - 1; // 图中第二种情况
    }
}
// @lc code=end
