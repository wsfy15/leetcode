/*
 * @lc app=leetcode.cn id=2397 lang=rust
 *
 * [2397] 被列覆盖的最多行数
 */

// @lc code=start
impl Solution {
    pub fn maximum_rows(matrix: Vec<Vec<i32>>, num_select: i32) -> i32 {
        let (n, m) = (matrix.len(), matrix[0].len());
        let mut masks = vec![0; n];
        for i in 0..n {
            for j in 0..m {
                masks[i] |= matrix[i][j] << j;
            }
        }

        let mut res = 0;
        for subset in 0..(1 << m) as i32 {
            if subset.count_ones() != num_select as u32 {
                continue;
            }

            let count = masks.iter().filter(|&&mask| mask & subset == mask).count();
            res = res.max(count);
        }

        res as i32
    }
}
// @lc code=end
