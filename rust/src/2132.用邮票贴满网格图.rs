/*
 * @lc app=leetcode.cn id=2132 lang=rust
 *
 * [2132] 用邮票贴满网格图
 */

// @lc code=start
impl Solution {
    pub fn possible_to_stamp(grid: Vec<Vec<i32>>, stamp_height: i32, stamp_width: i32) -> bool {
        let stamp_height = stamp_height as usize;
        let stamp_width = stamp_width as usize;

        let (n, m) = (grid.len(), grid[0].len());
        // 计算 grid 的二维前缀和
        let mut prefix = vec![vec![0; m + 1]; n + 1];
        for (i, row) in grid.iter().enumerate() {
            for (j, &val) in row.iter().enumerate() {
                prefix[i + 1][j + 1] = prefix[i][j + 1] + prefix[i + 1][j] - prefix[i][j] + val;
            }
        }

        // 计算二维差分
        let mut d = vec![vec![0; m + 2]; n + 2];
        for down in stamp_height as usize..=n {
            for right in stamp_width as usize..=m {
                let (left, up) = (right - stamp_width + 1, down - stamp_height + 1);
                // 检查邮票覆盖的区域内没有被占据的格子
                if prefix[down][right] - prefix[up - 1][right] - prefix[down][left - 1]
                    + prefix[up - 1][left - 1]
                    == 0
                {
                    // 将这个区域标记上邮票
                    d[up][left] += 1;
                    d[up][right + 1] -= 1;
                    d[down + 1][left] -= 1;
                    d[down + 1][right + 1] += 1;
                }
            }
        }

        //  还原二维差分矩阵对应的计数矩阵（原地计算）
        for (i, row) in grid.iter().enumerate() {
            for (j, &val) in row.iter().enumerate() {
                d[i + 1][j + 1] += d[i][j + 1] + d[i + 1][j] - d[i][j];
                if val == 0 && d[i + 1][j + 1] == 0 {
                    return false;
                }
            }
        }

        true
    }
}
// @lc code=end
