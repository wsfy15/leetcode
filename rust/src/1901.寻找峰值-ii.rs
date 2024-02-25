/*
 * @lc app=leetcode.cn id=1901 lang=rust
 *
 * [1901] 寻找峰值 II
 */

// @lc code=start
impl Solution {
    pub fn find_peak_grid(mat: Vec<Vec<i32>>) -> Vec<i32> {
        // 找到某一列的最大值和下标（行号）
        fn get_col_max(i: usize, mat: &Vec<Vec<i32>>) -> (usize, i32) {
            let m = mat.len();
            let mut index = 0;
            for j in 1..m {
                if mat[j][i] > mat[index][i] {
                    index = j;
                }
            }

            (index, mat[index][i])
        }

        fn get_val(i: i32, j: i32, mat: &Vec<Vec<i32>>) -> i32 {
            let (m, n) = (mat.len(), mat[0].len());
            if i < 0 || i as usize >= m || j < 0 || j as usize >= n {
                return -1;
            }
            mat[i as usize][j as usize]
        }

        let (mut left, mut right) = (0 as i32, (mat[0].len() - 1) as i32);
        while left < right {
            let mid = left + (right - left) / 2;
            let (max_index, max_val) = get_col_max(mid as usize, &mat);

            if max_val > get_val(max_index as i32, mid - 1, &mat)
                && max_val > get_val(max_index as i32, mid + 1, &mat)
            {
                return vec![max_index as i32, mid];
            }
            if get_val(max_index as i32, mid + 1, &mat) > max_val {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        let (idx, _) = get_col_max(left as usize, &mat);
        vec![idx as i32, left]
    }
}
// @lc code=end
