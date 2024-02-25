/*
 * @lc app=leetcode.cn id=2596 lang=rust
 *
 * [2596] 检查骑士巡视方案
 */

// @lc code=start
impl Solution {
    pub fn check_valid_grid(grid: Vec<Vec<i32>>) -> bool {
        let n = grid.len();
        let mut routes = vec![(0, 0); n * n];
        for i in 0..n {
            for j in 0..n {
                routes[grid[i][j] as usize] = (i, j);
            }
        }

        if routes[0] != (0, 0) {
            return false;
        }

        let (mut cur_row, mut cur_col) = (0, 0);
        let directions: Vec<(i32, i32)> = vec![
            (2, 1),
            (2, -1),
            (-2, 1),
            (-2, -1),
            (1, 2),
            (1, -2),
            (-1, 2),
            (-1, -2),
        ];
        for route in routes.iter().skip(1) {
            let (row, col) = (route.0 as i32, route.1 as i32);
            let mut success = false;
            for &dir in directions.iter() {
                if cur_row + dir.0 == row && cur_col + dir.1 == col {
                    cur_row = row;
                    cur_col = col;
                    success = true;
                    break;
                }
            }
            if !success {
                return false;
            }
        }

        true
    }
}
// @lc code=end
