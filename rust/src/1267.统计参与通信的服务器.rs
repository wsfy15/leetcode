/*
 * @lc app=leetcode.cn id=1267 lang=rust
 *
 * [1267] 统计参与通信的服务器
 */

// @lc code=start
impl Solution {
    pub fn count_servers(grid: Vec<Vec<i32>>) -> i32 {
        let (m, n) = (grid.len(), grid[0].len());
    
        let mut rows: Vec<i32> = vec![0; m];
        let mut cols: Vec<i32> = vec![0; n];
        for row in 0..m {
            for col in 0..n {
                if grid[row][col] == 1 {
                    rows[row] += 1;
                    cols[col] += 1;
                }
            }
        }
    
        let mut res = 0;
        for row in 0..m {
            for col in 0..n {
                if grid[row][col] == 1 && (rows[row] > 1 || cols[col] > 1) {
                    res += 1;
                }
            }
        }
    
        res
    }
}
// @lc code=end

